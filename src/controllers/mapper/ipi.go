package mapper

import "encoding/xml"

type IPI struct {
	XMLName xml.Name `xml:"IPI"`
	CEnq    string   `xml:"cEnq"`
	IPITrib IPITrib  `xml:"IPITrib"`
}
