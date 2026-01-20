package routes

import (
	"test-go/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine, handler *handlers.AuthHandler) {
	authGroup := router.Group("/auth")
	{
		// POST /auth/signup - Create a new user
		authGroup.POST("/signup", handler.SignUp)

		// POST /auth/login - Login user
		authGroup.POST("/login", handler.Login)
	}
}
