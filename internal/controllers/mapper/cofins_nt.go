package mapper

import "encoding/xml"

type COFINSNT struct {
	XMLName xml.Name `xml:"COFINSNT"`
	CST     string   `xml:"CST"`
}
