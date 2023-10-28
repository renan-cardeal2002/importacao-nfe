package xmlService

import (
	"importa-nfe/estruturas"
	"importa-nfe/conexao"
	"encoding/xml"
	"encoding/json"
	"log"
	"os"
	"fmt"
	"errors"
)

type ArquivoXML interface {
	Produtos() ([]byte, error)
	Emitente() ([]byte, error)
	Destinatario() ([]byte, error)
	InserirProdutos()
}

type NfeParser struct {
	data []byte
	nfe  estrut.NfeProc
}

func NewNfeParser(data []byte) (*NfeParser, error) {
	var nfe estrut.NfeProc
	err := json.Unmarshal(data, &nfe)
	if err != nil {
		return nil, err
	}

	return &NfeParser{
		data: data,
		nfe:  nfe,
	}, nil
}

func (np *NfeParser) Produtos() ([]byte, error) {
	produtos, err := json.Marshal(np.nfe.NFe.InfNFe.Det)
	if err != nil {
		return nil, err
	}
	return produtos, nil
}

func (np *NfeParser) Emitente() ([]byte, error) {
	emitente, err := json.Marshal(np.nfe.NFe.InfNFe.Emit)
	if err != nil {
		log.Fatal(err)
	}
	return emitente, nil
}

func (np *NfeParser) Destinatario() ([]byte, error) {
	destinatario, err := json.Marshal(np.nfe.NFe.InfNFe.Dest)
	if err != nil {
		log.Fatal(err)
	}
	return destinatario, nil
}

func VerificarEmit(cnpjEmit string, cnpjLogado string) error {
    db := database.ConexaoBd()
    defer db.Close()

    query := "SELECT cnpj FROM tbcadempresa WHERE cnpj = ?"
    var cnpj string

    err := db.QueryRow(query, cnpjLogado).Scan(&cnpj)
    if err != nil {
        return err
    }

    fmt.Println(cnpjEmit, cnpj)

    return nil
}

func InserirProdutos(produtosJSON string, cnpjEmit string) error {
    db := database.ConexaoBd()
    defer db.Close()

    var produtos []map[string]interface{}
    if err := json.Unmarshal([]byte(produtosJSON), &produtos); err != nil {
        return err
    }

    // Busque o ID da empresa emitente com base no CNPJ
    query := "SELECT id FROM tbcadempresa WHERE cnpj = ?"
    var empresaID int

    err := db.QueryRow(query, cnpjEmit).Scan(&empresaID)
    if err != nil {
        return err
    }

    insertStatement := "INSERT INTO tbcadprodutos (id_empresa, cProd, cEAN, xProd, uCom, qCom, vUnCom, vProd, vCusto, vPreco, vMargem, vAdicional) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

    for _, produto := range produtos {
        prod, ok := produto["Prod"].(map[string]interface{})
        if !ok {
            return nil
        }

        cEAN, cEANOK := prod["CEAN"].(string)

        // Verifica se o EAN já existe no banco de dados
        query := "SELECT count(1) as count FROM tbcadprodutos WHERE cEAN = ?"
        var countEan int

        err := db.QueryRow(query, cEAN).Scan(&countEan)
        if err != nil {
            return err
        }

        if countEan > 0 {
            return errors.New("EAN já existente")
        }

        cProd, cProdOK := prod["CProd"].(string)
        xProd, xProdOK := prod["XProd"].(string)
        uCom, uComOK := prod["UCom"].(string)
        qCom, qComOK := prod["QCom"].(string)
        vUnCom, vUnComOK := prod["VUnCom"].(string)

        vProd, vProdOK := prod["VProd"].(float64)
        vFrete := prod["VFrete"].(float64)
        vSeg := prod["VSeg"].(float64)
        vDesc := prod["VDesc"].(float64)
        vOutro := prod["VOutro"].(float64)

        vCusto := vProd + vFrete + vSeg + vOutro + vDesc
        vMargem := prod["VMargem"].(float64)

        vPreco := vCusto + vMargem
        vAdicional := prod["VAdicional"].(string)

        if cProdOK && cEANOK && xProdOK && uComOK && qComOK && vUnComOK && vProdOK {
            _, err := db.Exec(insertStatement, empresaID, cProd, cEAN, xProd, uCom, qCom, vUnCom, vProd, vCusto, vPreco, vMargem, vAdicional)
            if err != nil {
                return err
            }
        } else {
            return errors.New("Dados de produto inválidos")
        }
    }

    return nil
}

func InserirDest(destinatarioJSON string, cnpjEmit string) error {
    db := database.ConexaoBd()
    defer db.Close()

    var destinatario map[string]interface{}
    if err := json.Unmarshal([]byte(destinatarioJSON), &destinatario); err != nil {
        return err
    }
	
    // Busqa o ID da empresa destinatario com base no CNPJ
    query := "SELECT id FROM tbcadempresa WHERE cnpj = ?"
    var empresaID int

    err := db.QueryRow(query, cnpjEmit).Scan(&empresaID)
    if err != nil {
        return err
    }

    insertStatement := "INSERT INTO tbcadcliente (id_empresa, cnpj, xNome, email, xLgr, nro, xCpl, xBairro, cMun, CEP, fone) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

    cnpj, cnpjOK := destinatario["CNPJ"].(string)
    xNome, xNomeOK := destinatario["XNome"].(string)
    email, emailOK := destinatario["Email"].(string)


	enderDest, ok := destinatario["EnderDest"].(map[string]interface{})
    if !ok {
        return nil
    }

    xLgr, xLgrOK := enderDest["XLgr"].(string)
    nro, nroOK := enderDest["Nro"].(string)
    xCpl, xCplOK := enderDest["XCpl"].(string)
    xBairro, xBairroOK := enderDest["XBairro"].(string)
    cMun, cMunOK := enderDest["CMun"].(string)
    CEP, CEPOK := enderDest["CEP"].(string)
    fone, foneOK := enderDest["Fone"].(string)

	fmt.Println(cnpjOK, xNomeOK, emailOK, xLgrOK, nroOK, xCplOK, xBairroOK, cMunOK, CEPOK, foneOK)

    if cnpjOK && xNomeOK && emailOK && xLgrOK && nroOK && xCplOK && xBairroOK && cMunOK && CEPOK && foneOK {
        _, err := db.Exec(insertStatement, empresaID, cnpj, xNome, email, xLgr, nro, xCpl, xBairro, cMun, CEP, fone)
        if err != nil {
            return err
        }
    } else {
        return errors.New("Dados de destinatário inválidos")
    }

    return nil
}

func LerXml(xmlFile *os.File) ([]byte, error) {
	var nfeProc estrut.NfeProc

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
					return nil, err
				}
			}
		}
	}

	jsonData, err := json.Marshal(nfeProc)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
