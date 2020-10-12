package ports

import "github.com/unlar/alp-evaluator/internal/core/domain"

type PredictionService interface {
	Execute(img []uint8) (*domain.LicensePlate, error)
}
