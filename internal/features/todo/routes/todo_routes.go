package routes

import (
"test-go/internal/features/todo/handler"
"test-go/internal/shared/middleware"

"github.com/gin-gonic/gin"
)

func SetupTodoRoutes(router *gin.Engine, handler *handler.TodoHandler) {
router.Use(middleware.CORSMiddleware())

todoGroup := router.Group("/todos")
todoGroup.Use(middleware.AuthMiddleware())
{
todoGroup.POST("", handler.CreateTodo)
todoGroup.GET("", handler.GetAllTodos)
todoGroup.GET("/:id", handler.GetTodo)
todoGroup.PUT("/:id", handler.UpdateTodo)
todoGroup.DELETE("/:id", handler.DeleteTodo)
}
}
