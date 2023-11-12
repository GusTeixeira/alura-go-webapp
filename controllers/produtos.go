package controllers

import (
	"alura-go-webapp/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.GetProdutos()
	templates.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func InsertProduto(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)

		if err != nil {
			log.Println("Erro na conversão do preço")
		}

		quantidade, err := strconv.Atoi(r.FormValue("quantidade"))

		if err != nil {
			log.Println("Erro na conversão da quantidade")
		}

		models.CreateProduto(nome, descricao, preco, quantidade)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
