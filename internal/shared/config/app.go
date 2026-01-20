package config

import (
	"context"
	"net/http"
	"test-go/internal/features/auth/handler"
	authRoutes "test-go/internal/features/auth/routes"
	authUsecase "test-go/internal/features/auth/usecase"
	todoHandler "test-go/internal/features/todo/handler"
	"test-go/internal/features/todo/routes"
	todoUsecase "test-go/internal/features/todo/usecase"
	"test-go/internal/infrastructure/database/mongodb"
	infraRepository "test-go/internal/infrastructure/repository"
	"test-go/internal/shared/middleware"
	"test-go/internal/shared/response"
	"test-go/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router *gin.Engine
	Client *mongo.Client
}

type AppConfig struct {
	MongoURI       string
	DatabaseName   string
	CollectionName string
	JWTSecret      []byte
}

func NewApp(config *AppConfig) (*App, error) {
	appLogger := logger.InitLogger()

	client, err := mongodb.NewMongoClient(context.Background(), config.MongoURI)
	if err != nil {
		return nil, err
	}

	// Get database reference
	db := client.Database(config.DatabaseName)

	// Ensure indexes are created
	if err := mongodb.EnsureIndexes(context.Background(), db); err != nil {
		appLogger.WithError(err).Warn("Failed to create database indexes (non-fatal)")
	} else {
		appLogger.Info("Database indexes created successfully")
	}

	// Initialize repositories
	todoCollection := db.Collection(config.CollectionName)
	todoRepo := infraRepository.NewMongoTodoRepository(todoCollection)

	// Initialize todo use cases
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

	// Initialize auth repositories and use cases
	userCollection := db.Collection("users")
	userRepo := infraRepository.NewMongoUserRepository(userCollection)

	signUpUseCase := authUsecase.NewSignUpUseCase(userRepo)
	loginUseCase := authUsecase.NewLoginUseCase(userRepo)

	authHandler := handler.NewAuthHandler(signUpUseCase, loginUseCase, appLogger)

	router := gin.Default()

	// Add middleware
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.RequestLoggingMiddleware(appLogger))
	router.Use(middleware.RecoveryMiddleware(appLogger))

	// Swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check endpoint (no version required)
	router.GET("/health", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := client.Ping(ctx, nil); err != nil {
			appLogger.WithError(err).Error("database_health_check_failed")
			response.InternalServerError(c, "Database connection failed")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":    "healthy",
			"timestamp": time.Now().Unix(),
		})
	})

	// Setup versioned routes
	v1 := router.Group("/api/v1")
	{
		routes.SetupTodoRoutes(v1, todoHandlers, config.JWTSecret)
		authRoutes.SetupAuthRoutes(v1, authHandler)
	}

	return &App{
		Router: router,
		Client: client,
	}, nil
}
