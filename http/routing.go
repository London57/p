package http

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)


func NewRoute(g *gin.Engine, getall GetAllP) {

	if gin.Mode() != gin.ReleaseMode {
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	router := g.Group("/p")
	{
		router.GET("/getAll", getall.Exec)
	}
}