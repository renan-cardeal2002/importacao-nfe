package main

import (
    "github.com/gin-gonic/gin"
	"importa-nfe/controllers"
)

func main() {
    router := gin.Default()
    router.GET("/produtos", xmlController.GetProdutos)
    router.GET("/emitente", xmlController.GetEmitente)
    router.GET("/destinatario", xmlController.GetDestinatario)

    router.Run("localhost:8080")
}