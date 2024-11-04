package configuration

import "importa-nfe/internal/controllers"

type Handler struct {
	Importation controllers.ImportationController
}

func NewHandler() Handler {
	return Handler{
		Importation: controllers.NewImportationHandler(),
	}
}
