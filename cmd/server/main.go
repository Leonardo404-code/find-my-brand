package main

import (
	"net/http"
)

// @contact.name Leonardo de Farias Bispo
// @contact.email leonardobispo.dev@gmail.com
// @Title Find my brand
// @version 1.0
// @description Find my brand documentation
func main() {
	r := http.NewServeMux()

	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}
