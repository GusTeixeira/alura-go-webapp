package main

import (
	"log"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := conection()

	defer db.Close()

	log.Printf("Servidor iniciado na porta 8080")
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Println(err.Error())
	}

}

func index(w http.ResponseWriter, r *http.Request) {
	db := conection()

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

	templates.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
