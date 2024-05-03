package main

import (
	"log"
	"net/http"

	"github.com/Leonargo404-code/find-my-brand/pkg/brand/handler"
)

// @contact.name Leonardo de Farias Bispo
// @contact.email leonardobispo.dev@gmail.com
// @Title Find my brand
// @version 1.0
// @description Find my brand documentation
func main() {
	r := http.NewServeMux()
	brandHandlers := handler.Build()

	r.HandleFunc("/", brandHandlers.Search)

	log.Println("Server started at: http://localhost:3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}
