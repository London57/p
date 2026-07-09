package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/London57/gsqlc/datagen"
	_ "github.com/London57/gsqlc/docs"
	h "github.com/London57/gsqlc/http"
	"github.com/London57/gsqlc/postgres"
	"github.com/London57/gsqlc/server"
	"github.com/London57/gsqlc/service/p"
	tracer "github.com/London57/gsqlc/tracing"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func main() {
	// tracing
	tp := tracer.InitTracer()
	defer tp.Shutdown(context.Background())

	
	// db
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	ctx := context.Background()

	postgr := postgres.InitPostgresWithTracing(ctx, dsn)
	defer postgr.Close()

	// server
	srv := server.NewServ(":8080")
	srv.HttpServer.Handler = srv.App
	
	srv.App.Use(otelgin.Middleware("main-service", otelgin.WithTracerProvider(tp)))
	
	hand := h.GetAllP{}.New(
		p.GetAll{}.New(
			*datagen.New(postgr),
		),
	)
	h.NewRoute(srv.App, hand)

	err_ch := srv.Run()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)	
	
	select {
	case err := <- err_ch:
		fmt.Println(fmt.Errorf("srv error: %w", err))
	case err := <- interrupt:
		fmt.Println(fmt.Errorf("srv stopping: %s", err))
	}

	err := srv.Shutdown()
	if err != nil {
		fmt.Println(fmt.Errorf("srv error: %w", err))
	}
}