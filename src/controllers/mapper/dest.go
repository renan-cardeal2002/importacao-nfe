package mapper

import "encoding/xml"

type Dest struct {
	XMLName   xml.Name  `xml:"dest"`
	CNPJ      string    `xml:"CNPJ"`
	XNome     string    `xml:"xNome"`
	Email     string    `xml:"email"`
	EnderDest EnderDest `xml:"enderDest"`
}
