package mapper

import "encoding/xml"

type DetPag struct {
	XMLName xml.Name `xml:"detPag"`
	IndPag  string   `xml:"indPag"`
	TPag    string   `xml:"tPag"`
	VPag    float64  `xml:"vPag"`
}
