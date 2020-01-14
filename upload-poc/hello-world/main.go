package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Result struct {
	URL string `json:"url,omitempty"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	cfg := aws.Config{Region: aws.String("eu-west-1")}
	sess, err := session.NewSession(&cfg)
	if err != nil {
		panic("cannot establish AWS session")
	}

	contentType := request.QueryStringParameters["contentType"]
	name := request.QueryStringParameters["name"]
	bucket := os.Getenv("BUCKET")

	input := &s3.PutObjectInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(name),
		ContentType: aws.String(contentType),
	}

	svc := s3.New(sess)
	req, out := svc.PutObjectRequest(input)

	var url string
	if url, err = req.Presign(10 * time.Minute); err != nil {
		return createResponse("Cant presign url", 500)
	}
	result := &Result{URL: url}
	var b []byte
	if b, err = json.Marshal(result); err != nil {
		return createResponse("Cant marshal", 500)
	}
	return createResponse(string(b), 200)
}

func main() {
	lambda.Start(handler)
}

func createResponse(body string, statusCode int) (events.APIGatewayProxyResponse, error) {
	headers := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Headers": "Authorization,Content-Type",
		"Content-Type":                 "application/json",
	}
	return events.APIGatewayProxyResponse{
		Body:       body,
		Headers:    headers,
		StatusCode: statusCode,
	}, nil
}
