package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Serv struct {
	HttpServer *http.Server
	App *gin.Engine
}

func NewServ(addr string) *Serv {
	return &Serv{
		HttpServer: &http.Server{
			Addr: addr,
		},
		App: gin.Default(),
	}
}  

func (s *Serv) Run() <-chan error {
	ch := make(chan error, 1)
	go func() {
		ch <- s.HttpServer.ListenAndServe()
	}()
	return ch
}

func (s *Serv) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return s.HttpServer.Shutdown(ctx)
}
