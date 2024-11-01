package mapper

import "encoding/xml"

type ICMS struct {
	XMLName   xml.Name  `xml:"ICMS"`
	ICMSSN102 ICMSSN102 `xml:"ICMSSN102"`
}
