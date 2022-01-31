


package routes

import (
	"github.com/gin-gonic/gin"
	"go-gin-mongo-api/controllers"
)


func ProductRoutes(router *gin.Engine)  {
	router.POST("/product", controllers.CreateProduct)
	router.GET("/products", controllers.GetProducts)
}