package mapper

import "encoding/xml"

type InfNFe struct {
	XMLName    xml.Name   `xml:"infNFe"`
	Id         string     `xml:"Id,attr"`
	Versao     string     `xml:"versao,attr"`
	Ide        Ide        `xml:"ide"`
	Emit       Emit       `xml:"emit"`
	Dest       Dest       `xml:"dest"`
	Det        []Det      `xml:"det"`
	Total      Total      `xml:"total"`
	Transp     Transp     `xml:"transp"`
	Pag        Pag        `xml:"pag"`
	InfAdic    InfAdic    `xml:"infAdic"`
	InfRespTec InfRespTec `xml:"infRespTec"`
}
