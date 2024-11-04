package mapper

import "encoding/xml"

type ICMSSN102 struct {
	XMLName xml.Name `xml:"ICMSSN102"`
	Orig    string   `xml:"orig"`
	CSOSN   string   `xml:"CSOSN"`
}
