package main

import (
	"log"
	routes "todo-app/internal/routers"
	"todo-app/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	router := gin.Default()
	routes.SetupRoutes(router)

	log.Printf("Server is running on port 8080")
	log.Printf("API endpoints available at: http://localhost:8080/api/v1/todos")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
