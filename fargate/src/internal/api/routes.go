package api

import (
	"fmt"
	"net/http"
	"noter/src/internal/storage/dynamodb"

	"github.com/gorilla/mux"
)

func GetRouter(db *dynamodb.Adapter) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/healthy", healthyHandler)
	router.HandleFunc("/notes", HandleCreateNote(db)).Methods(http.MethodPost)
	router.HandleFunc("/notes", HandleListingNotes(db)).Methods(http.MethodGet)
	router.HandleFunc("/notes/{noteId}", HandleGetSingleNote(db)).Methods(http.MethodGet)
	return router
}

func healthyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "healthy\n")
}
