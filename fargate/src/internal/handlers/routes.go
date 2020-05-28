package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func API() *mux.Router {
	router := mux.NewRouter()
	n := Note{} //todo here goes db

	router.HandleFunc("/healthy", healthyHandler)
	router.HandleFunc("/notes", n.Create).Methods(http.MethodPost)
	router.HandleFunc("/notes", n.List).Methods(http.MethodGet)
	router.HandleFunc("/notes/{noteId:[0-9]+}", n.Get).Methods(http.MethodGet)
	router.HandleFunc("/notes/{noteId:[0-9]+}", n.Update).Methods(http.MethodPut)
	router.HandleFunc("/notes/{noteId:[0-9]+}", n.Delete).Methods(http.MethodDelete)
	return router
}

func healthyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "healthy\n")
}
