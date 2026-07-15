package http

import (
	http "github.com/London57/gsqlc/http/handlers"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


func NewRoute(g *gin.Engine, getall http.GetAllP, inserone http.InsertOneP) {

	if gin.Mode() != gin.ReleaseMode {
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	router := g.Group("/p")
	{
		router.GET("/getAll", getall.Exec)
		router.POST("/insertOne", inserone.Exec)
	}
}