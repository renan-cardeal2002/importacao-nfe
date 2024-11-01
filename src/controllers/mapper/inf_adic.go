package mapper

import "encoding/xml"

type InfAdic struct {
	XMLName xml.Name `xml:"infAdic"`
	InfCpl  string   `xml:"infCpl"`
}
