package ports

import (
	"github.com/unlar/alp-evaluator/internal/core/domain"
	"mime/multipart"
)

type PredictionService interface {
	Execute(file multipart.File, header *multipart.FileHeader) (*domain.LicensePlate, error)
}
