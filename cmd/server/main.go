package main

import (
	"test-go/internal/features/auth/usecase"
	"test-go/internal/shared/config"
	"test-go/internal/shared/middleware"
	"test-go/pkg/logger"
)

func main() {
	log := logger.InitLogger()

	config.LoadEnv()

	jwtSecret := config.GetEnv("JWT_SECRET", "H+a+b1XGVuSlkGpE8o3/h8aJFpbAk8Okvry0fluSzqs=")
	middleware.SetJWTSecret([]byte(jwtSecret))
	usecase.SetJWTSecretForUsecases([]byte(jwtSecret))

	port := config.GetEnv("PORT", "8080")
	mongoURI := config.GetEnv("MONGO_URI", "mongodb://localhost:27017")
	dbName := config.GetEnv("DATABASE_NAME", "todos")
	collName := config.GetEnv("COLLECTION_NAME", "todos")

	appConfig := &config.AppConfig{
		MongoURI:       mongoURI,
		DatabaseName:   dbName,
		CollectionName: collName,
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
