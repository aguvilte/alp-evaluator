package prediction

import (
	"net/http"
	"github.com/unlar/alp-evaluator/internal/core/ports"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/go-meli-toolkit/goutils/logger"
)

type Handler struct {
	service ports.PredictionService
}

func NewHandler(service ports.PredictionService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Execute(c *gin.Context) {

	mpFile, header, err := c.Request.FormFile("file")
	if err != nil {
		logger.Errorf("There was an error getting image from request.", err)
		c.String(http.StatusBadRequest, "")
	}

	plate, err := h.service.Execute(mpFile, header)
	if err != nil {
		c.String(http.StatusInternalServerError, "")
	}

	c.JSON(http.StatusOK, plate)
}