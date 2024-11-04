package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	database "importa-nfe/internal/infrastructure/configuration"
)

func InserirDest(destinatarioJSON string, cnpjEmit string) error {
	db := database.Connect()
	defer db.Close()

	var destinatario map[string]interface{}
	if err := json.Unmarshal([]byte(destinatarioJSON), &destinatario); err != nil {
		return err
	}

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
		return errors.New("dados de destinatário inválidos")
	}

	return nil
}
