package api

import (
	"context"
	"fmt"
	"net/http"
	"noter/src/internal/notes"
	"noter/src/internal/web"

	"github.com/gorilla/mux"
)

type notesGetter interface {
	GetSingleNote(ctx context.Context, noteId string) (*notes.Note, error)
	ListNotes(ctx context.Context) ([]*notes.Note, error)
}

type notePersiter interface {
	PersistNote(ctx context.Context, note *notes.Note) (*notes.Note, error)
}

func HandleCreateNote(persiter notePersiter) web.HttpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		var note notes.Note
		if err := web.Decode(r.Body, &note); err != nil {
			werr := fmt.Errorf("decode create note: %w", err)
			handleError(w, r, werr)
			return
		}
		fmt.Printf("note %v", note)
		persistedNote, err := persiter.PersistNote(r.Context(), &note)
		fmt.Printf("persisted note %v", note)
		if err != nil {
			handleError(w, r, err)
			return
		}

		web.Respond(w, r, persistedNote, http.StatusOK)
	}
}

func HandleGetSingleNote(getter notesGetter) web.HttpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Printf("request %v\n", r)
		noteId, ok := vars["noteId"]
		if !ok {
			werr := fmt.Errorf("note id not found")
			handleError(w, r, werr)
			return
		}
		fmt.Printf("noteId %v\n", noteId)

		note, err := getter.GetSingleNote(r.Context(), noteId)
		fmt.Printf("notes %v\n", note)
		if err != nil {
			handleError(w, r, err)
			return
		}

		web.Respond(w, r, note, http.StatusOK)
	}
}

func HandleListingNotes(getter notesGetter) web.HttpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		notes, err := getter.ListNotes(r.Context())
		fmt.Printf("notes %v\n", notes)
		if err != nil {
			handleError(w, r, err)
			return
		}

		web.Respond(w, r, notes, http.StatusOK)
	}
}

//todo improve error handling
func handleError(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Printf("Error: %v", err)
	web.Respond(w, r, err, http.StatusInternalServerError)
}
