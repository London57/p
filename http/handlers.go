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

func (h *GetAllP) Exec(c *gin.Context) {
	ps, err := h.service.Exec(c)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, ps)
}