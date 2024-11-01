package main

import (
	"github.com/gin-gonic/gin"
	"importa-nfe/src/controllers"
)

func main() {
	router := gin.Default()

	router.GET("/inserirNfe", controllers.InserirNFE)
	router.GET("/produtos", controllers.GetProdutos)
	router.GET("/emitente", controllers.GetEmitente)
	router.GET("/destinatario", controllers.GetDestinatario)

	err := router.Run("localhost:2000")
	if err != nil {
		return
	}
}
