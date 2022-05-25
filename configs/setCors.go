
package configs

import (
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetCors() gin.HandlerFunc{

	return cors.New(cors.Config{
		AllowOrigins:     []string{"http://loaclhost:3000"},
		AllowMethods:     []string{"POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://loaclhost:3000"
		},
		MaxAge: 12 * time.Hour,
	})

}