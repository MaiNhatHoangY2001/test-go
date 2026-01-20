package routes

import (
"test-go/internal/features/auth/handler"

"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine, handler *handler.AuthHandler) {
authGroup := router.Group("/auth")
{
authGroup.POST("/signup", handler.SignUp)
authGroup.POST("/login", handler.Login)
}
}
