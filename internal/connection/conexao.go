package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	dsn := "myuser:mypassword@tcp(localhost:3306)/importanfe"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("erro ao realizar conexão ao banco:", err)
		return nil
	}

	if err = db.Ping(); err != nil {
		fmt.Println("erro ao realizar conexão ao banco:", err)
		return nil
	}

	return db
}
