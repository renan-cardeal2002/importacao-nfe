package main

import (
	"encoding/xml"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// crio os tipos de dados da NFe
type NfeProc struct {
    XMLName xml.Name `xml:"nfeProc"`
    Versao  string   `xml:"versao,attr"`
    NFe     NFe      `xml:"NFe"`
    ProtNFe ProtNFe  `xml:"protNFe"`
}

type NFe struct {
    XMLName xml.Name `xml:"NFe"`
    InfNFe  InfNFe   `xml:"infNFe"`
}

type InfNFe struct {
    XMLName xml.Name `xml:"infNFe"`
    Id      string   `xml:"Id,attr"`
    Versao  string   `xml:"versao,attr"`
    Ide     Ide      `xml:"ide"`
    Emit    Emit     `xml:"emit"`
    Dest    Dest     `xml:"dest"`
    Det     []Det    `xml:"det"`
    Total   Total    `xml:"total"`
    Transp  Transp   `xml:"transp"`
    Pag     Pag      `xml:"pag"`
    InfAdic InfAdic  `xml:"infAdic"`
    InfRespTec InfRespTec `xml:"infRespTec"`
}

type Ide struct {
    XMLName xml.Name `xml:"ide"`
    CUf     string   `xml:"cUF"`
    CNF     string   `xml:"cNF"`
    NatOp   string   `xml:"natOp"`
    // Adicione os campos restantes aqui.
}

type Emit struct {
    XMLName xml.Name `xml:"emit"`
    CNPJ    string   `xml:"CNPJ"`
}

type Dest struct {
    XMLName   xml.Name `xml:"dest"`
    CNPJ      string   `xml:"CNPJ"`
    XNome     string   `xml:"xNome"`
    Email     string   `xml:"email"`
    EnderDest EnderDest `xml:"enderDest"`
}

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

type Det struct {
    XMLName xml.Name `xml:"det"`
    NItem   string   `xml:"nItem,attr"`
    Prod    Prod     `xml:"prod"`
    Imposto Imposto  `xml:"imposto"`
}

type Prod struct {
    XMLName xml.Name `xml:"prod"`
    CProd   string   `xml:"cProd"`
    CEAN    string   `xml:"cEAN"`
    XProd   string   `xml:"xProd"`
    UCom    string   `xml:"uCom"`
    QCom    string   `xml:"qCom"`
    VUnCom  string   `xml:"vUnCom"`
    VFrete  float64  `xml:"vFrete"`
    VSeg    float64  `xml:"vSeg"`
    VDesc   float64  `xml:"vDesc"`
    VOutro  float64  `xml:"vOutro"`
    VCusto  float64
    VPreco  float64
}

// vMargem e vAdicional - não achei no xml

type Imposto struct {
    XMLName  xml.Name `xml:"imposto"`
    VTotTrib float64  `xml:"vTotTrib"`
    ICMS     ICMS     `xml:"ICMS"`
    IPI      IPI      `xml:"IPI"`
    PIS      PIS      `xml:"PIS"`
    COFINS   COFINS   `xml:"COFINS"`
}

type ICMS struct {
    XMLName xml.Name `xml:"ICMS"`
    ICMSSN102 ICMSSN102 `xml:"ICMSSN102"`
}

type ICMSSN102 struct {
    XMLName xml.Name `xml:"ICMSSN102"`
    Orig string `xml:"orig"`
    CSOSN string `xml:"CSOSN"`
}

type IPI struct {
    XMLName xml.Name `xml:"IPI"`
    CEnq   string `xml:"cEnq"`
    IPITrib IPITrib `xml:"IPITrib"`
}

type IPITrib struct {
    XMLName xml.Name `xml:"IPITrib"`
    CST string `xml:"CST"`
    VBC float64 `xml:"vBC"`
    PIPI float64 `xml:"pIPI"`
    VIPI float64 `xml:"vIPI"`
}

type PIS struct {
    XMLName xml.Name `xml:"PIS"`
    PISNT PISNT `xml:"PISNT"`
}

type PISNT struct {
    XMLName xml.Name `xml:"PISNT"`
    CST string `xml:"CST"`
}

type COFINS struct {
    XMLName xml.Name `xml:"COFINS"`
    COFINSNT COFINSNT `xml:"COFINSNT"`
}

type COFINSNT struct {
    XMLName xml.Name `xml:"COFINSNT"`
    CST string `xml:"CST"`
}

type Total struct {
    XMLName xml.Name `xml:"total"`
    ICMSTot ICMSTot `xml:"ICMSTot"`
}

type ICMSTot struct {
    XMLName xml.Name `xml:"ICMSTot"`
    VBC         float64 `xml:"vBC"`
    VICMS       float64 `xml:"vICMS"`
    VICMSDeson  float64 `xml:"vICMSDeson"`
    VFCP        float64 `xml:"vFCP"`
    // Adicione os campos restantes aqui.
}

type Transp struct {
    XMLName xml.Name `xml:"transp"`
    ModFrete string `xml:"modFrete"`
    // Adicione os campos restantes aqui.
}

type Pag struct {
    XMLName xml.Name `xml:"pag"`
    DetPag DetPag `xml:"detPag"`
}

type DetPag struct {
    XMLName xml.Name `xml:"detPag"`
    IndPag string `xml:"indPag"`
    TPag string `xml:"tPag"`
    VPag float64 `xml:"vPag"`
}

type InfAdic struct {
    XMLName xml.Name `xml:"infAdic"`
    InfCpl string `xml:"infCpl"`
}

type InfRespTec struct {
    XMLName xml.Name `xml:"infRespTec"`
    CNPJ string `xml:"CNPJ"`
    XContato string `xml:"xContato"`
    Email string `xml:"email"`
    Fone string `xml:"fone"`
}

type ProtNFe struct {
    XMLName xml.Name `xml:"protNFe"`
    Versao  string   `xml:"versao,attr"`
    InfProt InfProt  `xml:"infProt"`
}

type InfProt struct {
    XMLName xml.Name `xml:"infProt"`
    TpAmb   string   `xml:"tpAmb"`
    VerAplic string   `xml:"verAplic"`
    ChNFe   string   `xml:"chNFe"`
    DhRecbto string   `xml:"dhRecbto"`
    NProt   string   `xml:"nProt"`
    DigVal  string   `xml:"digVal"`
    CStat   string   `xml:"cStat"`
    XMotivo string   `xml:"xMotivo"`
}

// leio o arquivo xml e retorno os dados da nfe
func lerXml() []byte {
	xmlFile, err := os.Open("./docs/41230910541434000152550010000012411749316397-nfe.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer xmlFile.Close()

	var nfeProc NfeProc
	
	decoder := xml.NewDecoder(xmlFile)

	for {
		token, err := decoder.Token()
		if err != nil {
			break
		}

		switch se := token.(type) {
		case xml.StartElement:
			if se.Name.Local == "nfeProc" {
				if err := decoder.DecodeElement(&nfeProc, &se); err != nil {
					log.Fatal(err)
				}
			}
		}
	}

	jsonData, err := json.Marshal(nfeProc)
	if err != nil {
		log.Fatal(err)
	}

	return jsonData
}

func emitente(data []byte) []byte {
    var nfe NfeProc
    if err := json.Unmarshal(data, &nfe); err != nil {
        log.Fatal(err)
    }

    emitente, err := json.Marshal(nfe.NFe.InfNFe.Emit)
	if err != nil {
		log.Fatal(err)
	}

    return emitente
}

func produtos(data []byte) []byte {
    var nfe NfeProc
    if err := json.Unmarshal(data, &nfe); err != nil {
        log.Fatal(err)
    }

    produtos, err := json.Marshal(nfe.NFe.InfNFe.Det)
	if err != nil {
		log.Fatal(err)
	}
    
    return produtos
}

func main() {
    data := lerXml()

    emit := emitente(data)
    prod := produtos(data)

    fmt.Println(string(emit))
    fmt.Println(string(prod))
}