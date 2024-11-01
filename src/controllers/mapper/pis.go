package mapper

import "encoding/xml"

type PIS struct {
	XMLName xml.Name `xml:"PIS"`
	PISNT   PISNT    `xml:"PISNT"`
}
