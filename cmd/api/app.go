package api

import (
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApp() {
	Dependencies.Initialize()
	mapRoutes()
	run("8080")
}

func run(port string) {
	router.Run(":" + port)
}

func mapRoutes() {
	router.GET("/ping", Dependencies.PingHandler.Ping)
}