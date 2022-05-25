package main

import (
	"go-gin-mongo-api/configs"
	"go-gin-mongo-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default();
	// disable trust all proxies 
	app.SetTrustedProxies(nil)
	app.Use(configs.SetCors())
	apiVersionOne := app.Group("/api/v1")
    routes.ProductRoutes(apiVersionOne)
	app.Run(configs.ServerConfigurations.Port)	
}