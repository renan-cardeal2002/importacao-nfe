package pkg

import (
	"encoding/json"
	"encoding/xml"
	"importa-nfe/internal/core/domain"
	"importa-nfe/internal/infrastructure/adapter/inbound/controllers/mapper"
	"os"
)

func ReadXML(xmlFile *os.File) ([]byte, error) {
	var nfeProc mapper.NfeProc

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

func NewNfeParser(data []byte) (domain.NfeParser, error) {
	var nfe mapper.NfeProc
	err := json.Unmarshal(data, &nfe)
	if err != nil {
		return domain.NfeParser{}, err
	}

	return domain.NfeParser{
		Data: data,
		Nfe:  nfe,
	}, nil
}
