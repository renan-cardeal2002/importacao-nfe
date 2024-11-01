package mapper

import "encoding/xml"

type NfeProc struct {
	XMLName xml.Name `xml:"nfeProc"`
	Versao  string   `xml:"versao,attr"`
	NFe     NFe      `xml:"NFe"`
	ProtNFe ProtNFe  `xml:"protNFe"`
}
