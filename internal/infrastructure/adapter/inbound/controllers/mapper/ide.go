package mapper

import "encoding/xml"

type Ide struct {
	XMLName xml.Name `xml:"ide"`
	CUf     string   `xml:"cUF"`
	CNF     string   `xml:"cNF"`
	NatOp   string   `xml:"natOp"`
}
