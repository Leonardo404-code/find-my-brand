package main

import (
	"log"
	"net/http"

	"github.com/Leonargo404-code/find-my-brand/pkg/brand/handler"
	"github.com/rs/cors"
)

// @contact.name Leonardo de Farias Bispo
// @contact.email leonardobispo.dev@gmail.com
// @Title Find my brand
// @version 1.0
// @description Find my brand documentation
func main() {
	r := http.NewServeMux()

	fs := http.FileServer(http.Dir("./assets"))
	r.Handle("/", fs)

	brandHandlers := handler.Build()
	r.HandleFunc("/search", brandHandlers.Search)

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodPost},
	})

	handler := cors.Handler(r)

	log.Println("Server started at: http://localhost:3000")
	if err := http.ListenAndServe(":3000", handler); err != nil {
		panic(err)
	}
}
