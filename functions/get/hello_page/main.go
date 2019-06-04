package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	html = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <h2>Hello World!</h2>
</body>
</html>
	`
)

// Define request and response
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	res := Response{
		StatusCode: 200,
		Body:       html,
		Headers: map[string]string{
			"Content-Type": "text/html",
		},
	}

	return res, nil
}

func main() {
	lambda.Start(Handler)
}
