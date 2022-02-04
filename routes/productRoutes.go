package routes

import (
	"go-gin-mongo-api/controllers"
	"github.com/gin-gonic/gin"
)


func ProductRoutes(routerVersion *gin.RouterGroup)  {
	
	{	
		routerVersion.POST("/product", controllers.CreateProduct)
		routerVersion.GET("/products", controllers.GetProducts)
		routerVersion.GET("/product/:productId", controllers.GetProduct)
		routerVersion.PUT("/update-product", controllers.UpdateProduct);
		routerVersion.DELETE("/product/:productId", controllers.UpdateProduct);
	}

}