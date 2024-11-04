package xmlService

import (
	"encoding/json"
	"importa-nfe/internal/core/domain"
	"log"
)

type ArquivoXML interface {
	Produtos() ([]byte, error)
	Emitente() ([]byte, error)
	Destinatario() ([]byte, error)
	InserirProdutos()
}

func Produtos(np domain.NfeParser) ([]byte, error) {
	produtos, err := json.Marshal(np.Nfe.NFe.InfNFe.Det)
	if err != nil {
		return nil, err
	}
	return produtos, nil
}

func Emitente(np domain.NfeParser) ([]byte, error) {
	emitente, err := json.Marshal(np.Nfe.NFe.InfNFe.Emit)
	if err != nil {
		log.Fatal(err)
	}
	return emitente, nil
}

func Destinatario(np domain.NfeParser) ([]byte, error) {
	destinatario, err := json.Marshal(np.Nfe.NFe.InfNFe.Dest)
	if err != nil {
		log.Fatal(err)
	}
	return destinatario, nil
}
