package api

import (
	"github.com/unlar/alp-evaluator/internal/core/ports"
	"github.com/unlar/alp-evaluator/internal/core/services"
	"github.com/unlar/alp-evaluator/internal/handlers/ping"
	"github.com/unlar/alp-evaluator/internal/handlers/prediction"
	"github.com/unlar/alp-evaluator/internal/repositories/rest"
	"github.com/unlar/alp-evaluator/internal/repositories/os"
)

var dependencies = &DependenciesDefinitions{}

type DependenciesDefinitions struct {
	//repositories
	RestRepository ports.RestRepository
	OSRepository   ports.OSRepository

	//services
	PredictionService ports.PredictionService

	//handlers
	PingHandler       *ping.Handler
	PredictionHandler *prediction.Handler
}

func (d *DependenciesDefinitions) Initialize() {
	d.RestRepository = rest.NewProductionRepo()
	d.OSRepository = os.NewRepo()

	d.PredictionService = services.NewPredictionService(d.RestRepository, d.OSRepository)

	d.PingHandler = ping.NewHandler()
	d.PredictionHandler = prediction.NewHandler(d.PredictionService)
}
