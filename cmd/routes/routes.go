package routes

import (
	"github.com/gin-gonic/gin"
	"importa-nfe/internal/infrastructure/configuration"
)

func Setup(controller configuration.Handler) *gin.Engine {
	router := gin.Default()

	router.GET("/inserirNfe", controller.Importation.InserirNFE)
	router.GET("/produtos", controller.Importation.GetProdutos)
	router.GET("/emitente", controller.Importation.GetEmitente)
	router.GET("/destinatario", controller.Importation.GetDestinatario)

	return router
}
