package repository

import (
	"fmt"
	database "importa-nfe/internal/infrastructure/configuration"
)

func VerificarEmit(cnpjEmit string, cnpjLogado string) error {
	db := database.Connect()
	defer db.Close()

	query := "SELECT cnpj FROM tbcadempresa WHERE cnpj = ?"
	var cnpj string

	err := db.QueryRow(query, cnpjLogado).Scan(&cnpj)
	if err != nil {
		return err
	}

	fmt.Println(cnpjEmit, cnpj)

	return nil
}
