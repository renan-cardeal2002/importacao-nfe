package main

import (
	"encoding/xml"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"importa-nfe/estruturas"
)

type ArquivoXML interface {
    produtos(data []byte) ([]byte, error)
}

// leio o arquivo xml e retorno os dados da nfe
func lerXml() []byte {
	xmlFile, err := os.Open("./docs/41230910541434000152550010000012411749316397-nfe.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer xmlFile.Close()

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
					log.Fatal(err)
				}
			}
		}
	}

	jsonData, err := json.Marshal(nfeProc)
	if err != nil {
		log.Fatal(err)
	}

	return jsonData
}

func emitente(data []byte) []byte {
    var nfe estrut.NfeProc
	
    if err := json.Unmarshal(data, &nfe); err != nil {
        log.Fatal(err)
    }

    emitente, err := json.Marshal(nfe.NFe.InfNFe.Emit)
	if err != nil {
		log.Fatal(err)
	}

    return emitente
}

func produtos(data []byte) []byte {
    var nfe estrut.NfeProc
	
    if err := json.Unmarshal(data, &nfe); err != nil {
		log.Fatal(err)
    }

    produtos, err := json.Marshal(nfe.NFe.InfNFe.Det)
	if err != nil {
		log.Fatal(err)
	}
    
    return produtos
}

func main() {
    data := lerXml()

    emit := emitente(data)
    prod := produtos(data)

    fmt.Println(string(emit))
    fmt.Println(string(prod))
}