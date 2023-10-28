// database/database.go
package database

import (
    "database/sql"
	_ "github.com/go-sql-driver/mysql"
    "log"
)

// ConexaoBd estabelece e retorna uma conexão com o banco de dados
func ConexaoBd() *sql.DB {
    // Configurar informações de conexão com o banco de dados
    dsn := "root@tcp(localhost:3306)/importanfe"

    // Abrir uma conexão com o banco de dados
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    if err = db.Ping(); err != nil {
        log.Fatal(err)
    }
    return db
}
