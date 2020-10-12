package api

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApp() {
	dependencies.Initialize()
	mapRoutes()
	run("8080")
}

func run(port string) {
	router.Run(":" + port)
}

func mapRoutes() {
	router.GET("/ping", dependencies.PingHandler.Ping)
	router.POST("/recognize/license_plate", dependencies.PredictionHandler.Execute)
}
