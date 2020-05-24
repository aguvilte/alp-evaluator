package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {}

func NewHandler() *Handler {
	return &Handler{}
}

const pong = "pong"

func (h *Handler) Ping(c *gin.Context) {
	c.String(http.StatusOK, pong)
}