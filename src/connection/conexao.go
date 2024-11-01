package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	dsn := "root@tcp(localhost:3306)/importanfe"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		_ = fmt.Errorf("erro ao realizar conexão ao banco")
	}

	if err = db.Ping(); err != nil {
		_ = fmt.Errorf("erro ao realizar conexão ao banco")
	}

	return db
}
