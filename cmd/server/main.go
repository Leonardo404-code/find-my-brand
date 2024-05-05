package main

import (
	"log"
	"net/http"
	"os"

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

	log.Printf("Server started at: http://localhost:%s", os.Getenv("PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), handler); err != nil {
		panic(err)
	}
}
