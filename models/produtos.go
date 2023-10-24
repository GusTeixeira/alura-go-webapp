package models

import (
	"alura-go-webapp/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetProdutos() []Produto {
	db := db.Connection()

	query, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for query.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := query.Scan(&id, &nome, &descricao, &quantidade, &preco)

		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}
