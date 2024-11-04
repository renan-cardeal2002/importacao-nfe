package mapper

import "encoding/xml"

type Total struct {
	XMLName xml.Name `xml:"total"`
	ICMSTot ICMSTot  `xml:"ICMSTot"`
}
