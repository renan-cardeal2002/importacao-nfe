package ports

import (
	"context"
	"importa-nfe/internal/core/domain"
)

type ProdutosRepository interface {
	InserirProdutos(produtosJSON string, CNPJ string) error
}

type DestinatarioRepository interface {
	InserirDest(ctx context.Context, EmpresaID int, destinatarioJSON string) error
}

type EmpresaRepository interface {
	FindByCNPJ(ctx context.Context, CNPJ string) (domain.Empresa, error)
}
