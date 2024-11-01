package mapper

import "encoding/xml"

type PISNT struct {
	XMLName xml.Name `xml:"PISNT"`
	CST     string   `xml:"CST"`
}
