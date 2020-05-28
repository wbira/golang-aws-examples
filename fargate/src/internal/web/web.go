package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request)

func Respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
	w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			fmt.Printf("Error %v when encoding: %b", err, data)
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
	}
	w.WriteHeader(status)
}
