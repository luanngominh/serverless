package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Define request and response
type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request Request) (Response, error) {
	fmt.Println("Received body: ", request.Body)

	return Response{
		StatusCode: 200,
		Body:       request.Body,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
