package mapper

import "encoding/xml"

type Pag struct {
	XMLName xml.Name `xml:"pag"`
	DetPag  DetPag   `xml:"detPag"`
}
