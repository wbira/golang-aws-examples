package handlers

import (
	"net/http"
	"noter/src/internal/web"
)

type Note struct {
	Id   int
	Text string
}

func (n *Note) Create(w http.ResponseWriter, r *http.Request) {

}

func (n *Note) Get(w http.ResponseWriter, r *http.Request) {

}

func (n *Note) List(w http.ResponseWriter, r *http.Request) {
	notes := []*Note{
		{Id: 1, Text: "Serverless for beginners"},
		{Id: 2, Text: "Fargate how to strat"},
	}

	web.Respond(w, r, notes, http.StatusOK)
}

func (n *Note) Update(w http.ResponseWriter, r *http.Request) {

}

func (n *Note) Delete(w http.ResponseWriter, r *http.Request) {

}
