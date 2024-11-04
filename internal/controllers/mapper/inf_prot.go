package mapper

import "encoding/xml"

type InfProt struct {
	XMLName  xml.Name `xml:"infProt"`
	TpAmb    string   `xml:"tpAmb"`
	VerAplic string   `xml:"verAplic"`
	ChNFe    string   `xml:"chNFe"`
	DhRecbto string   `xml:"dhRecbto"`
	NProt    string   `xml:"nProt"`
	DigVal   string   `xml:"digVal"`
	CStat    string   `xml:"cStat"`
	XMotivo  string   `xml:"xMotivo"`
}
