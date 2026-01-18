package routes

import (
	"todo-app/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	todoController := controllers.NewTodoController()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"status":  "healthy",
		})
	})

	v1 := router.Group("/api/v1")
	{
		todos := v1.Group("/todos")
		{
			todos.GET("", todoController.GetAllTodos)
			todos.POST("", todoController.CreateTodo)
			todos.GET("/:id", todoController.GetTodoByID)
			todos.PUT("/:id", todoController.UpdateTodo)
			todos.DELETE("/:id", todoController.DeleteTodo)
		}
	}
}
