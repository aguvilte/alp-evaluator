package services

import (
	"github.com/unlar/alp-evaluator/internal/core/domain"
	"github.com/unlar/alp-evaluator/internal/core/ports"
)

type predictionService struct {
	repository ports.Repository
}

func NewPredictionService(repository ports.Repository) *predictionService {
	return &predictionService{repository: repository}
}

func (p *predictionService) Execute(img []uint8) (*domain.LicensePlate, error) {
	plate, err := p.repository.GetPrediction(img)
	if err != nil {
		return nil, err
	}
	return plate, nil
}
