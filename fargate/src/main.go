package main

import (
	"net/http"
	"noter/src/internal/handlers"
)

func main() {
	api := handlers.API()
	http.Handle("/", api)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("Internal server error")
	}
}
