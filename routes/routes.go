package routes

import (
	"alura-go-webapp/controllers"
	"net/http"
)

func LoadRotes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.InsertProduto)
}
