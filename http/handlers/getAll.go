package http

import (
	"net/http"

	"github.com/London57/gsqlc/service/p"
	"github.com/gin-gonic/gin"
)

type GetAllP struct {
	service pservice.GetAll
}

func (GetAllP) New(service pservice.GetAll) GetAllP {
	return GetAllP{
		service: service,
	}
}

// @Summary Get All Ps
// @Tags P
// @Accept json
// @Produce json
// @Success 200 {object} []datagen.P
// @Success 200 {string} string "no p yet" "Empty list"
// @Failure 500 {object} map[string]interface{} "Internal Server Error"
// @Router /p/getAll [get]
func (h *GetAllP) Exec(c *gin.Context) {
	ps, err := h.service.Exec(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if ps == nil {
		c.JSON(http.StatusOK, "no p yet")
		return
	}
	c.JSON(http.StatusOK, ps)
}