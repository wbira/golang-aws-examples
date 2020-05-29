package api

import (
	"fmt"
	"net/http"
	"noter/src/internal/storage"

	"github.com/gorilla/mux"
)

func GetRouter(repository *storage.Repository) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/healthy", healthyHandler)
	router.HandleFunc("/notes", HandleCreateNote(repository)).Methods(http.MethodPost)
	router.HandleFunc("/notes", HandleListingNotes(repository)).Methods(http.MethodGet)
	router.HandleFunc("/notes/{noteId}", HandleGetSingleNote(repository)).Methods(http.MethodGet)
	return router
}

func healthyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "healthy\n")
}
