package mapper

import "encoding/xml"

type Emit struct {
	XMLName xml.Name `xml:"emit"`
	CNPJ    string   `xml:"CNPJ"`
}
