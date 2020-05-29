package main

import (
	"fmt"
	"net/http"
	"noter/src/internal/api"
	"noter/src/internal/storage"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	sess, err := session.NewSession()
	if err != nil {
		wErr := fmt.Errorf("cannot create AWS session, error: %w", err)
		panic(wErr)
	}
	dbCli := dynamodb.New(sess)
	repository := storage.NewRepository(dbCli, "Noter") //todo fetch table name from env variable
	http.Handle("/", api.GetRouter(repository))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("Internal server error")
	}
}
