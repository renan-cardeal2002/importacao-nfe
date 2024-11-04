package repository

import (
	"context"
	"database/sql"
	"importa-nfe/internal/core/domain"
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

func (r empresaRepository) FindByCNPJ(ctx context.Context, CNPJ string) (domain.Empresa, error) {
	query := "SELECT id, cnpj FROM empresa WHERE cnpj = ?"

	var empresa domain.Empresa

	err := r.db.QueryRowContext(ctx, query, CNPJ).Scan(&empresa.ID, &empresa.CNPJ)
	if err != nil {
		return domain.Empresa{}, err
	}

	return empresa, nil
}
