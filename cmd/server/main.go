package main

import (
	"test-go/internal/features/auth/usecase"
	"test-go/internal/shared/config"
	"test-go/pkg/logger"

	_ "test-go/docs" // Swagger docs
)

// @title Todo API
// @version 1.0
// @description This is a production-ready Todo API with user authentication and pagination
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@todoapi.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	log := logger.InitLogger()

	config.LoadEnv()

	jwtSecret := config.GetEnv("JWT_SECRET", "")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}
	
	jwtSecretBytes := []byte(jwtSecret)
	usecase.SetJWTSecretForUsecases(jwtSecretBytes)

	port := config.GetEnv("PORT", "8080")
	mongoURI := config.GetEnv("MONGO_URI", "mongodb://localhost:27017")
	dbName := config.GetEnv("DATABASE_NAME", "todos")
	collName := config.GetEnv("COLLECTION_NAME", "todos")

	appConfig := &config.AppConfig{
		MongoURI:       mongoURI,
		DatabaseName:   dbName,
		CollectionName: collName,
		JWTSecret:      jwtSecretBytes,
	}

	app, err := config.NewApp(appConfig)
	if err != nil {
		log.Fatal("Failed to initialize app:", err)
	}

	log.WithField("port", port).Info("server listening")

	if err := app.Router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
