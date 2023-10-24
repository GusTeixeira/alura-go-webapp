package main

import (
	"alura-go-webapp/routes"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	log.Printf("Servidor iniciado na porta 8000")
	routes.LoadRotes()
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Println(err.Error())
	}

}
