package main

import (
    "github.com/gin-gonic/gin"
	"importa-nfe/controllers"
    "importa-nfe/conexao"
)

func main() {
    db := database.ConexaoBd()
	
    router := gin.Default()
    router.GET("/produtos", xmlController.GetProdutos)
    router.GET("/emitente", xmlController.GetEmitente)
    router.GET("/destinatario", xmlController.GetDestinatario)

    router.Run("localhost:8080")

	defer db.Close()
}
