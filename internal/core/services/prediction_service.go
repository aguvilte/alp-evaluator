package services

import (
	"github.com/unlar/alp-evaluator/internal/core/domain"
	"github.com/unlar/alp-evaluator/internal/core/ports"
	"mime/multipart"
)

type predictionService struct {
	restRepo ports.RestRepository
	osRepo   ports.OSRepository
}

func NewPredictionService(restRepo ports.RestRepository, osRepo ports.OSRepository) *predictionService {
	return &predictionService{restRepo: restRepo, osRepo: osRepo}
}

func (p *predictionService) Execute(file multipart.File, header *multipart.FileHeader) (*domain.LicensePlate, error) {
	err := p.osRepo.SaveFile(file, header)

	plate, err := p.restRepo.GetPrediction(file, header)

	err = p.osRepo.DeleteFile(header.Filename)

	return plate, err
}
