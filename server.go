package main

import (
	// "fmt"
	// "net/http"
	"go-gin-mongo-api/configs"
	"github.com/gin-gonic/gin"
	"go-gin-mongo-api/routes"
)


func main() {
	router := gin.Default();
	serverConfigs := configs.SetServerConfigurations();
	// connect to database
	configs.ConnectToMongDb();

	//use products router
    routes.ProductRoutes(router)

	// start  the API server
	router.Run(serverConfigs.Port);
	
}