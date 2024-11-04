package configuration

import (
	"importa-nfe/internal/infrastructure/adapter/inbound/controllers"
)

type Handler struct {
	Importation controllers.ImportationController
}

func NewHandler(port *Port) Handler {
	return Handler{
		Importation: controllers.NewImportationHandler(port.produtosRepository),
	}
}
