package main

import (
	"test-go/internal/features/auth/usecase"
	"test-go/internal/shared/config"
	"test-go/pkg/logger"
)

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
