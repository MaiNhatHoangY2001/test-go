package config

import (
	"context"
	"test-go/internal/api/handlers"
	"test-go/internal/api/routes"
	"test-go/internal/application/usecases"
	"test-go/internal/infrastructure/database/mongodb"
	"test-go/internal/infrastructure/repository"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

type AppConfig struct {
	MongoURI       string
	DatabaseName   string
	CollectionName string
}

func NewApp(config *AppConfig) (*App, error) {
	client, err := mongodb.NewMongoClient(context.Background(), config.MongoURI)
	if err != nil {
		return nil, err
	}
	collection := client.Database(config.DatabaseName).Collection(config.CollectionName)
	repo := repository.NewMongoTodoRepository(collection)

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
	}, nil
}
