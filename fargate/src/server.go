package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	Id   int
	Name string
}

var books = []*Book{{Id: 1, Name: "Serverless for beginners"}, {Id: 2, Name: "Fargate how to strat"}}

func main() {
	http.HandleFunc("/", HealthyHandler)
	http.HandleFunc("/books", GetBooks)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("Internal server error")
	}
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	respond(w, r, books, http.StatusOK)
}

func HealthyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Healthy\n")
}

func respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.Header().Set("Access-Control-Allow-Headers", "Authorization,Content-Type")

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			fmt.Printf("Error %v when encoding: %b", err, data)
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
	}
	w.WriteHeader(status)
}
