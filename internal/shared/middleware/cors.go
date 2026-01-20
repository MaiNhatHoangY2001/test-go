package middleware

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	
	// Get allowed origins from environment variable
	allowedOriginsEnv := os.Getenv("ALLOWED_ORIGINS")
	if allowedOriginsEnv != "" {
		config.AllowOrigins = strings.Split(allowedOriginsEnv, ",")
	} else {
		// Default to localhost for development
		config.AllowOrigins = []string{"http://localhost:3000", "http://localhost:8080"}
	}
	
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowCredentials = true

	return cors.New(config)
}
