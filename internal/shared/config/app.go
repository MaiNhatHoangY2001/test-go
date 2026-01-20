package config

import (
	"context"
	"test-go/internal/features/auth/handler"
	authRoutes "test-go/internal/features/auth/routes"
	authUsecase "test-go/internal/features/auth/usecase"
	todoHandler "test-go/internal/features/todo/handler"
	"test-go/internal/features/todo/routes"
	todoUsecase "test-go/internal/features/todo/usecase"
	"test-go/internal/infrastructure/database/mongodb"
	infraRepository "test-go/internal/infrastructure/repository"
	"test-go/internal/shared/middleware"
	"test-go/pkg/logger"

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
	appLogger := logger.InitLogger()

	client, err := mongodb.NewMongoClient(context.Background(), config.MongoURI)
	if err != nil {
		return nil, err
	}

	todoCollection := client.Database(config.DatabaseName).Collection(config.CollectionName)
	todoRepo := infraRepository.NewMongoTodoRepository(todoCollection)

	createUseCase := todoUsecase.NewCreateTodoUseCase(todoRepo)
	getTodoUseCase := todoUsecase.NewGetTodoUseCase(todoRepo)
	getAllUseCase := todoUsecase.NewGetAllTodosUseCase(todoRepo)
	updateUseCase := todoUsecase.NewUpdateTodoUseCase(todoRepo)
	deleteUseCase := todoUsecase.NewDeleteTodoUseCase(todoRepo)

	todoHandlers := todoHandler.NewTodoHandler(
		createUseCase,
		getTodoUseCase,
		getAllUseCase,
		updateUseCase,
		deleteUseCase,
		appLogger,
	)

	userCollection := client.Database(config.DatabaseName).Collection("users")
	userRepo := infraRepository.NewMongoUserRepository(userCollection)

	signUpUseCase := authUsecase.NewSignUpUseCase(userRepo)
	loginUseCase := authUsecase.NewLoginUseCase(userRepo)

	authHandler := handler.NewAuthHandler(signUpUseCase, loginUseCase, appLogger)

	router := gin.Default()

	// Add middleware
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.RequestLoggingMiddleware(appLogger))
	router.Use(middleware.RecoveryMiddleware(appLogger))

	routes.SetupTodoRoutes(router, todoHandlers)
	authRoutes.SetupAuthRoutes(router, authHandler)

	return &App{
		Router: router,
	}, nil
}
