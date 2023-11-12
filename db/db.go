package db

import (
	"database/sql"
)

func Connection() (db *sql.DB) {
	conexao := "user=postgres dbname=alura_loja password=300402 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
