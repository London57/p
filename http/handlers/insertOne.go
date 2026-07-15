package http

import (
	"net/http"

	db "github.com/London57/gsqlc/datagen"
	"github.com/London57/gsqlc/service/p"
	"github.com/gin-gonic/gin"
)

type InsertOneP struct {
	service pservice.InsertOne
}

func (InsertOneP) New(service pservice.InsertOne) InsertOneP {
	return InsertOneP{
		service: service,
	}
}

// @Summary Insert One P
// @Tags P
// @Accept json
// @Produce json
// @Success 201 {object} datagen.P "Record created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /p/insertOne [post]
func (h *InsertOneP) Exec(c *gin.Context) {
	var args db.InsertOnePParams
	if err := c.ShouldBindJSON(args); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ps, err := h.service.Exec(c.Request.Context(), args)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, ps)
}