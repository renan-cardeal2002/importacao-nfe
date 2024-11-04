package mapper

import "encoding/xml"

type ICMSTot struct {
	XMLName    xml.Name `xml:"ICMSTot"`
	VBC        float64  `xml:"vBC"`
	VICMS      float64  `xml:"vICMS"`
	VICMSDeson float64  `xml:"vICMSDeson"`
	VFCP       float64  `xml:"vFCP"`
}
