package routes

import (
	"test-go/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupTodoRoutes(router *gin.Engine, handler *handlers.TodoHandler) {
	todoGroup := router.Group("/todos")
	{
		//POST /todos - Create a new todo
		todoGroup.POST("", handler.CreateTodo)

		//GET /todos - Get all todos
		todoGroup.GET("", handler.GetAllTodos)

		// GET /todos/:id - Get specific todo
		todoGroup.GET("/:id", handler.GetTodo)

		// PUT /todos/:id - Update a todo
		todoGroup.PUT("/:id", handler.UpdateTodo)

		// DELETE /todos/:id - Delete a todo
		todoGroup.DELETE("/:id", handler.DeleteTodo)
	}
}
