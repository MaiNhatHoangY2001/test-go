package routes

import (
	"test-go/internal/features/todo/handler"
	"test-go/internal/shared/middleware"

	"github.com/gin-gonic/gin"
)

func SetupTodoRoutes(router gin.IRouter, handler *handler.TodoHandler, jwtSecret []byte) {
	todoGroup := router.Group("/todos")
	todoGroup.Use(middleware.AuthMiddleware(jwtSecret))
	{
		todoGroup.POST("", handler.CreateTodo)
		todoGroup.GET("", handler.GetAllTodos)
		todoGroup.GET("/:id", handler.GetTodo)
		todoGroup.PUT("/:id", handler.UpdateTodo)
		todoGroup.DELETE("/:id", handler.DeleteTodo)
	}
}
