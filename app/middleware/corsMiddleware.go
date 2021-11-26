package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type corsMiddleware struct {
}

var CorsMiddleware = new(corsMiddleware)

func (corsMiddleware *corsMiddleware) Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	config.ExposeHeaders = []string{"New-Token", "New-Expires-In", "Content-Disposition"}

	return cors.New(config)
}
