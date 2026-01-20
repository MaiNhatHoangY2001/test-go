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
	todoCollection := client.Database(config.DatabaseName).Collection(config.CollectionName)
	todoRepo := repository.NewMongoTodoRepository(todoCollection)

	userCollection := client.Database(config.DatabaseName).Collection("users")
	userRepo := repository.NewMongoUserRepository(userCollection)

	createUseCase := usecases.NewCreateTodoUseCase(todoRepo)
	getTodoUseCase := usecases.NewGetTodoUseCase(todoRepo)
	getAllUseCase := usecases.NewGetAllTodosUseCase(todoRepo)
	updateUseCase := usecases.NewUpdateTodoUseCase(todoRepo)
	deleteUseCase := usecases.NewDeleteTodoUseCase(todoRepo)

	signUpUseCase := usecases.NewSignUpUseCase(userRepo)
	loginUseCase := usecases.NewLoginUseCase(userRepo)

	todoHandlers := handlers.NewTodoHandler(
		createUseCase,
		getTodoUseCase,
		getAllUseCase,
		updateUseCase,
		deleteUseCase,
	)

	authHandler := handlers.NewAuthHandler(signUpUseCase, loginUseCase)

	router := gin.Default()

	routes.SetupTodoRoutes(router, todoHandlers)
	routes.SetupAuthRoutes(router, authHandler)

	return &App{
		Router: router,
	}, nil
}
