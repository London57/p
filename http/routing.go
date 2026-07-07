package http

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)


func NewRoute(g *gin.Engine, getall GetAllP) {
	tracing := otelgin.Middleware("main service")
	g.Use(tracing)
	router := g.Group("p")
	{
		router.GET("/getAll", getall.Exec)
	}
}