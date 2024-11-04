package ports

import (
	"context"
	"database/sql"
	"importa-nfe/internal/core/domain"
)

type ProdutosRepository interface {
	FindProdutosByEmpresaID(ctx context.Context, EmpresaID int) ([]domain.Produto, error)
	InserirProdutos(ctx context.Context, produtosJSON string, EmpresaID int) error
}

type DestinatarioRepository interface {
	InserirDest(ctx context.Context, EmpresaID int, destinatarioJSON string) error
}

type EmpresaRepository interface {
	FindByCNPJ(ctx context.Context, CNPJ string) (domain.Empresa, error)
}

type TransactionManager interface {
	Begin(ctx context.Context) (*sql.Tx, error)
	Commit(tx *sql.Tx) error
	Rollback(tx *sql.Tx) error
}
