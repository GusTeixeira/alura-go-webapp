package routes

import (
	"alura-go-webapp/controllers"
	"net/http"
)

func LoadRotes() {
	http.HandleFunc("/", controllers.Index)
}
