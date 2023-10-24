package controllers

import (
	"alura-go-webapp/models"
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.GetProdutos()
	templates.ExecuteTemplate(w, "Index", todosOsProdutos)
}
