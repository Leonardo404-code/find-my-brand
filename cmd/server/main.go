package main

import (
	"net/http"
)

func main() {
	r := http.NewServeMux()

	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}
