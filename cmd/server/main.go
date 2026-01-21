package main

import (
	"test-go/internal/features/auth/usecase"
	"test-go/internal/shared/config"
	"test-go/pkg/logger"

	_ "test-go/docs" // Swagger docs
)

// gin-swagger middleware
// swagger embed files

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
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
