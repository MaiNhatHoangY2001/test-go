package main

import (
	"log"
	"test-go/internal/config"
)

func main() {
	config.LoadEnv()

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

	// Start server on port 8080
	if err := app.Router.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
