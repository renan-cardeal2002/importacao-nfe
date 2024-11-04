package mapper

import "encoding/xml"

type InfRespTec struct {
	XMLName  xml.Name `xml:"infRespTec"`
	CNPJ     string   `xml:"CNPJ"`
	XContato string   `xml:"xContato"`
	Email    string   `xml:"email"`
	Fone     string   `xml:"fone"`
}
