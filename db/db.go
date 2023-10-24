package db

import (
	"database/sql"
)

func conection() (db *sql.DB) {
	conexao := "user=postgres dbname=alura_loja password=Guga3004_ host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
