package mapper

import "encoding/xml"

type Det struct {
	XMLName xml.Name `xml:"det"`
	NItem   string   `xml:"nItem,attr"`
	Prod    Prod     `xml:"prod"`
	Imposto Imposto  `xml:"imposto"`
}
