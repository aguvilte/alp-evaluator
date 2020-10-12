package ports

import (
	"github.com/unlar/alp-evaluator/internal/core/domain"
)

type Repository interface {
	GetPrediction(img []uint8) (*domain.LicensePlate, error)
}
