package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

const pong = "pong"

func (h *Handler) Ping(c *gin.Context) {
	c.String(http.StatusOK, pong)
}
