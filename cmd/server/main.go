package main

import (
	"log"
	"test-go/internal/config"
)

func main() {
	app := config.NewApp()

	if err := app.Router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
