package config

import (
	"test-go/internal/api/handlers"
	"test-go/internal/api/routes"
	"test-go/internal/application/usecases"
	"test-go/internal/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

func NewApp() *App {
	repo := repository.NewInMemoryTodoRepository()

	createUseCase := usecases.NewCreateTodoUseCase(repo)
	getTodoUseCase := usecases.NewGetTodoUseCase(repo)
	getAllUseCase := usecases.NewGetAllTodosUseCase(repo)
	updateUseCase := usecases.NewUpdateTodoUseCase(repo)
	deleteUseCase := usecases.NewDeleteTodoUseCase(repo)

	handlers := handlers.NewTodoHandler(
		createUseCase,
		getTodoUseCase,
		getAllUseCase,
		updateUseCase,
		deleteUseCase,
	)

	router := gin.Default()

	routes.SetupTodoRoutes(router, handlers)

	return &App{
		Router: router,
	}
}
