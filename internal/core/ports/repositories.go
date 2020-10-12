package ports

import (
	"github.com/unlar/alp-evaluator/internal/core/domain"
	"mime/multipart"
)

type RestRepository interface {
	GetPrediction(file multipart.File, header *multipart.FileHeader) (*domain.LicensePlate, error)
}

type OSRepository interface {
	SaveFile(file multipart.File, header *multipart.FileHeader) error
	DeleteFile(filename string) error
}