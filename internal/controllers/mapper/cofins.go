package mapper

import "encoding/xml"

type COFINS struct {
	XMLName  xml.Name `xml:"COFINS"`
	COFINSNT COFINSNT `xml:"COFINSNT"`
}
