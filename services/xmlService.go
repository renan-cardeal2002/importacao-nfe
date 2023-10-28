package xmlService

import (
	"encoding/xml"
	"encoding/json"
	"log"
	"importa-nfe/estruturas"
	"os"
)

type ArquivoXML interface {
	Produtos() ([]byte, error)
	Emitente() ([]byte, error)
	Destinatario() ([]byte, error)
}

type NfeParser struct {
	data []byte
	nfe  estrut.NfeProc
}

func NewNfeParser(data []byte) (*NfeParser, error) {
	var nfe estrut.NfeProc
	err := json.Unmarshal(data, &nfe)
	if err != nil {
		return nil, err
	}

	return &NfeParser{
		data: data,
		nfe:  nfe,
	}, nil
}

func (np *NfeParser) Produtos() ([]byte, error) {
	produtos, err := json.Marshal(np.nfe.NFe.InfNFe.Det)
	if err != nil {
		return nil, err
	}
	return produtos, nil
}

func (np *NfeParser) Emitente() ([]byte, error) {
	emitente, err := json.Marshal(np.nfe.NFe.InfNFe.Emit)
	if err != nil {
		log.Fatal(err)
	}
	return emitente, nil
}

func (np *NfeParser) Destinatario() ([]byte, error) {
	destinatario, err := json.Marshal(np.nfe.NFe.InfNFe.Dest)
	if err != nil {
		log.Fatal(err)
	}
	return destinatario, nil
}

func LerXml(xmlFile *os.File) ([]byte, error) {
	var nfeProc estrut.NfeProc

	decoder := xml.NewDecoder(xmlFile)

	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if se.Name.Local == "nfeProc" {
				if err := decoder.DecodeElement(&nfeProc, &se); err != nil {
					return nil, err
				}
			}
		}
	}

	jsonData, err := json.Marshal(nfeProc)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
