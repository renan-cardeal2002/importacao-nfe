package mapper

import "encoding/xml"

type NFe struct {
	XMLName xml.Name `xml:"NFe"`
	InfNFe  InfNFe   `xml:"infNFe"`
}
