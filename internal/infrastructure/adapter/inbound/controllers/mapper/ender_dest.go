package mapper

import "encoding/xml"

type EnderDest struct {
	XMLName xml.Name `xml:"enderDest"`
	XLgr    string   `xml:"xLgr"`
	Nro     string   `xml:"nro"`
	XCpl    string   `xml:"xCpl"`
	XBairro string   `xml:"xBairro"`
	CMun    string   `xml:"cMun"`
	CEP     string   `xml:"CEP"`
	Fone    string   `xml:"fone"`
}
