package main

import (
	"net/http"
	"noter/src/internal/api"
	"noter/src/internal/storage/dynamodb"
)

func main() {
	db := &dynamodb.Adapter{}
	http.Handle("/", api.GetRouter(db))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("Internal server error")
	}
}
