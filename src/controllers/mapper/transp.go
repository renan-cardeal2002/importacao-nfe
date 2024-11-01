package mapper

import "encoding/xml"

type Transp struct {
	XMLName  xml.Name `xml:"transp"`
	ModFrete string   `xml:"modFrete"`
	// Adicione os campos restantes aqui.
}
