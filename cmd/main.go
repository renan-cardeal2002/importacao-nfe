package main

import (
	"importa-nfe/cmd/routes"
	"importa-nfe/internal/infrastructure/configuration"
	"log"
)

func main() {
	log.Println("Starting application...")

	port, err := configuration.NewPort()
	if err != nil {
		log.Fatalf("error on init ports: %s", err)
	}

	handler := configuration.NewHandler(port)
	router := routes.Setup(handler)

	err = router.Run("localhost:8088")
	if err != nil {
		return
	}
}
