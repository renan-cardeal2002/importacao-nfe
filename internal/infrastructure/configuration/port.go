package configuration

import (
	"fmt"
	"importa-nfe/internal/core/ports"
	"importa-nfe/internal/infrastructure/adapter/outbound/repository"
)

type Port struct {
	produtosRepository     ports.ProdutosRepository
	destinatarioRepository ports.DestinatarioRepository
	empresaRepository      ports.EmpresaRepository
	transactionManager     ports.TransactionManager
}

func NewPort() (*Port, error) {

	db, err := newMySqlDB()
	if err != nil {
		return nil, fmt.Errorf("init database connection: %w", err)
	}

	produtosRepository := repository.NewProdutosRepository(db)
	destinatarioRepository := repository.NewDestinatarioRepository(db)
	empresaRepository := repository.NewEmpresaRepository(db)
	transactionManager := repository.NewTransactionManager(db)

	return &Port{
		produtosRepository:     produtosRepository,
		destinatarioRepository: destinatarioRepository,
		empresaRepository:      empresaRepository,
		transactionManager:     transactionManager,
	}, nil
}
