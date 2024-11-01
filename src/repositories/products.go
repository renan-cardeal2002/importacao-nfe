package repositories

import (
	"encoding/json"
	"errors"
	database "importa-nfe/src/connection"
)

func InserirProdutos(produtosJSON string, cnpjEmit string) error {
	db := database.Connect()
	defer db.Close()

	var produtos []map[string]interface{}
	if err := json.Unmarshal([]byte(produtosJSON), &produtos); err != nil {
		return err
	}

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

		vMargem := 0.0 //prod["VMargem"].(float64)

		vPreco := vCusto + vMargem
		vAdicional := 0.0 //prod["VAdicional"].(string)

		if cProdOK && cEANOK && xProdOK && uComOK && qComOK && vUnComOK && vProdOK {
			_, err := db.Exec(insertStatement, empresaID, cProd, cEAN, xProd, uCom, qCom, vUnCom, vProd, vCusto, vPreco, vMargem, vAdicional)
			if err != nil {
				return err
			}
		} else {
			return errors.New("dados de produto inválidos")
		}
	}

	return nil
}
