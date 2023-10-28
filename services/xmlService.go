package xmlService

import (
	"encoding/xml"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"importa-nfe/estruturas"
)

type ArquivoXML interface {
	Produtos() ([]byte, error)
	Emitente() ([]byte, error)
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

func main() {
	xmlFile, err := os.Open("./docs/41230910541434000152550010000012411749316397-nfe.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer xmlFile.Close()

	data, err := LerXml(xmlFile)
	if err != nil {
		log.Fatal(err)
	}

	parser, err := NewNfeParser(data)
	if err != nil {
		log.Fatal(err)
	}

	emit, err := parser.Emitente()
	if err != nil {
		log.Fatal(err)
	}

	prod, err := parser.Produtos()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(emit))
	fmt.Println(string(prod))
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
