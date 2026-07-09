package http

import (
	"net/http"

	"github.com/London57/gsqlc/service/p"
	"github.com/gin-gonic/gin"
)

type GetAllP struct {
	service p.GetAll
}

func (GetAllP) New(service p.GetAll) GetAllP {
	return GetAllP{
		service: service,
	}
}

// @Summary Get All Ps
// @Tags P
// @Accept json
// @Produce json
// @Success 200 {object} []datagen.P
// @Failure 500 {string} string "Internal Server Error"
// @Router /p/getAll [get]
func (h *GetAllP) Exec(c *gin.Context) {
	ps, err := h.service.Exec(c.Request.Context())
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	if ps == nil {
		c.JSON(http.StatusOK, "no p yet")
		return
	}
	c.JSON(http.StatusOK, ps)
}