package dynamodb

import (
	"context"
	"noter/src/internal/notes"
)

type Adapter struct {
}

var list = []notes.Note{
	{
		NoteId: "214e2wdq",
		Text:   "Some note",
	},
}

func (d *Adapter) PersistNote(ctx context.Context, note notes.Note) (notes.Note, error) {
	list = append(list, note)
	return note, nil
}

func (d *Adapter) GetSingleNote(ctx context.Context, noteId string) (notes.Note, error) {
	for _, note := range list {
		if note.NoteId == noteId {
			return note, nil
		}
	}
	return notes.Note{}, nil
}

func (d *Adapter) ListNotes(ctx context.Context) ([]notes.Note, error) {
	return list, nil
}
