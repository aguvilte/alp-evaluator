package api

import "github.com/unlar/alp-evaluator/internal/handlers/ping"

func (d *DependenciesDefinitions) initProductionEnv() {
	d.PingHandler = ping.NewHandler()
}
