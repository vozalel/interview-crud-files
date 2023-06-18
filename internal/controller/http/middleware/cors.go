package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"PUT", "GET", "DELETE"},
		AllowHeaders:    []string{"Authorization, X-Requested-With, Access-Control-Request-Headers, Access-Control-Allow-Headers, Access-Control-Request-Method, Origin, Accept, Content-Type"},

		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
