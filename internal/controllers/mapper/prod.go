package mapper

import "encoding/xml"

type Prod struct {
	XMLName xml.Name `xml:"prod"`
	CProd   string   `xml:"cProd"`
	CEAN    string   `xml:"cEAN"`
	XProd   string   `xml:"xProd"`
	UCom    string   `xml:"uCom"`
	QCom    string   `xml:"qCom"`
	VUnCom  string   `xml:"vUnCom"`
	VProd   float64  `xml:"vProd"`
	VFrete  float64  `xml:"vFrete"`
	VSeg    float64  `xml:"vSeg"`
	VDesc   float64  `xml:"vDesc"`
	VOutro  float64  `xml:"vOutro"`
	VCusto  float64
	VPreco  float64
}
