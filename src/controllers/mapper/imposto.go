package mapper

import "encoding/xml"

type Imposto struct {
	XMLName  xml.Name `xml:"imposto"`
	VTotTrib float64  `xml:"vTotTrib"`
	ICMS     ICMS     `xml:"ICMS"`
	IPI      IPI      `xml:"IPI"`
	PIS      PIS      `xml:"PIS"`
	COFINS   COFINS   `xml:"COFINS"`
}
