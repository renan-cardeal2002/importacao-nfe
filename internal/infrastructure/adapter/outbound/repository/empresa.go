package repository

import (
	"database/sql"
	"fmt"
	"importa-nfe/internal/core/ports"
)

type empresaRepository struct {
	db *sql.DB
}

func NewEmpresaRepository(db *sql.DB) ports.EmpresaRepository {
	return empresaRepository{
		db: db,
	}
}

func (r empresaRepository) VerificarEmit(cnpjEmit string, cnpjLogado string) error {
	query := "SELECT cnpj FROM tbcadempresa WHERE cnpj = ?"
	var cnpj string

	err := r.db.QueryRow(query, cnpjLogado).Scan(&cnpj)
	if err != nil {
		return err
	}

	fmt.Println(cnpjEmit, cnpj)

	return nil
}
