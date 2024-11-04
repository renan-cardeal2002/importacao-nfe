package repository

import (
	"context"
	"database/sql"
	"importa-nfe/internal/core/ports"
)

type transactionManager struct {
	db *sql.DB
}

func NewTransactionManager(db *sql.DB) ports.TransactionManager {
	return &transactionManager{
		db: db,
	}
}

func (tm *transactionManager) Begin(ctx context.Context) (*sql.Tx, error) {
	return tm.db.BeginTx(ctx, nil)
}

func (tm *transactionManager) Commit(tx *sql.Tx) error {
	return tx.Commit()
}

func (tm *transactionManager) Rollback(tx *sql.Tx) error {
	return tx.Rollback()
}
