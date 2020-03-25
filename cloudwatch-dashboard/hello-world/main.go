package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// handlers returns errors to generate logs on dashboard
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Info: recived request %v", request)

	if value, ok := request.QueryStringParameters["error"]; ok {
		if value == "notfound" {
			log.Print("Error: not found")
			return createApiGatewayResponse("Not found\n", http.StatusNotFound)
		}
		log.Print("Error: internal server error")
		return createApiGatewayResponse("Error\n", http.StatusInternalServerError)
	}

	return createApiGatewayResponse("Success\n", http.StatusOK)
}

func createApiGatewayResponse(body string, statusCode int) (events.APIGatewayProxyResponse, error) {
	var err error

	if statusCode == http.StatusInternalServerError {
		err = errors.New("internal server error")
	}
	return events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: statusCode,
	}, err
}

func main() {
	lambda.Start(handler)
}
