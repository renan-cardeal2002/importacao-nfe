package mapper

import "encoding/xml"

type IPITrib struct {
	XMLName xml.Name `xml:"IPITrib"`
	CST     string   `xml:"CST"`
	VBC     float64  `xml:"vBC"`
	PIPI    float64  `xml:"pIPI"`
	VIPI    float64  `xml:"vIPI"`
}
