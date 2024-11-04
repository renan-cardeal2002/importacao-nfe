package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"importa-nfe/internal/core/ports"
)

type destinatarioRepository struct {
	db *sql.DB
}

func NewDestinatarioRepository(db *sql.DB) ports.DestinatarioRepository {
	return destinatarioRepository{
		db: db,
	}
}

func (r destinatarioRepository) InserirDest(ctx context.Context, EmpresaID int, destinatarioJSON string) error {
	query := `
		INSERT INTO clientes 
		    (id_empresa, cnpj, x_nome, email, x_lgr, nro, x_cpl, x_bairro, c_mun, cep, fone) 
		VALUES 
		    ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	var destinatario map[string]interface{}
	if err := json.Unmarshal([]byte(destinatarioJSON), &destinatario); err != nil {
		return err
	}

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
		_, err := r.db.ExecContext(ctx, query, EmpresaID, cnpj, xNome, email, xLgr, nro, xCpl, xBairro, cMun, CEP, fone)
		if err != nil {
			return err
		}
	} else {
		return errors.New("dados de destinatário inválidos")
	}

	return nil
}
