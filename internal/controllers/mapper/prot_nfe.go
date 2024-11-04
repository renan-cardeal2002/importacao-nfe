package mapper

import "encoding/xml"

type ProtNFe struct {
	XMLName xml.Name `xml:"protNFe"`
	Versao  string   `xml:"versao,attr"`
	InfProt InfProt  `xml:"infProt"`
}
