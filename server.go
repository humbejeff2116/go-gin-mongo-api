package main

import (
	"go-gin-mongo-api/configs"
	"go-gin-mongo-api/routes"

	"github.com/gin-gonic/gin"
)


func main() {
	app := gin.Default();
	serverConfigs := configs.SetServerConfigurations();
	// connect to database
	configs.ConnectToMongDb();
	// disable trust all proxies for now as no proxy client has been used to make request to server
	app.SetTrustedProxies(nil)
	// enable cors
	app.Use(configs.SetCors())
	// use logger middleware
	app.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	app.Use(gin.Recovery())
	// api/v1 routes group
	apiVersionOne := app.Group("/api/v1")
	// products routes
    routes.ProductRoutes(apiVersionOne)
	// users routes
	routes.UsersRoutes(apiVersionOne)
	// start  the API server
	app.Run(serverConfigs.Port)	
}