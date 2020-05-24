package api

import "github.com/unlar/alp-evaluator/internal/handlers/ping"

var Dependencies = &DependenciesDefinitions{}

type DependenciesDefinitions struct {
	PingHandler		*ping.Handler
}

func (d *DependenciesDefinitions) Initialize() {
	d.initProductionEnv()
}