package main

import (
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Info: recived request %v", request)

	if _, ok := request.QueryStringParameters["error"]; ok {
		log.Print("Error: internal server error")
		return createApiGatewayResponse("Error\n", http.StatusInternalServerError)
	}

	return createApiGatewayResponse("Success\n", http.StatusOK)
}

func createApiGatewayResponse(body string, statusCode int) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: statusCode,
	}, nil
}

func main() {
	lambda.Start(handler)
}
