package repositories

import (
	"database/sql"
	"encoding/json"
	"errors"
	database "importa-nfe/internal/connection"
)

func InserirProdutos(produtosJSON string, cnpjEmit string) error {
	db := database.Connect()
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

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

		cEAN := prod["CEAN"].(string)

		query = "SELECT count(1) as count FROM tbcadprodutos WHERE cEAN = ?"
		var countEan int

		err = db.QueryRow(query, cEAN).Scan(&countEan)
		if err != nil {
			return err
		}

		if countEan > 0 {
			return errors.New("EAN já existente")
		}

		cProd := prod["CProd"].(string)
		xProd := prod["XProd"].(string)
		uCom := prod["UCom"].(string)
		qCom := prod["QCom"].(string)
		vUnCom := prod["VUnCom"].(string)

		vProd := prod["VProd"].(float64)
		vFrete := prod["VFrete"].(float64)
		vSeg := prod["VSeg"].(float64)
		vDesc := prod["VDesc"].(float64)
		vOutro := prod["VOutro"].(float64)

		vCusto := vProd + vFrete + vSeg + vOutro + vDesc

		vMargem := 0.0 // VMargem não está presente no xml

		vPreco := vCusto + vMargem
		vAdicional := 0.0 // VMargem não está presente no xml

		_, err = db.Exec(insertStatement, empresaID, cProd, cEAN, xProd, uCom, qCom, vUnCom, vProd, vCusto, vPreco, vMargem, vAdicional)
		if err != nil {
			return err
		}
	}

	return nil
}
