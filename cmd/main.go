package main

import (
	"importa-nfe/cmd/routes"
	"importa-nfe/internal/infrastructure/configuration"
	"log"
)

func main() {
	log.Println("Starting application...")
	handler := configuration.NewHandler()
	router := routes.Setup(handler)

	err := router.Run("localhost:8088")
	if err != nil {
		return
	}
}
