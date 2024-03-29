package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	tableName = ""
)

func init() {
	tableName = os.Getenv("DYNAMODB_TABLE")
}

// Define request and response
type Response events.APIGatewayProxyResponse

type contact struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)

	result, err := svc.Scan(&dynamodb.ScanInput{
		TableName: aws.String("meocon_contacts"),
	})
	if err != nil {
		return Response{
			StatusCode: 404,
			Body:       fmt.Sprintf("Query error with messages %v", err),
		}, nil
	}

	contacts := []contact{}
	for _, item := range result.Items {
		c := contact{
			Name:    *item["name"].S,
			Email:   *item["email"].S,
			Message: *item["message"].S,
		}
		contacts = append(contacts, c)
	}

	data, err := json.Marshal(contacts)
	if err != nil {
		return Response{
			StatusCode: 404,
			Body:       fmt.Sprintf("Decode json error %v", err),
		}, nil
	}

	res := Response{
		StatusCode: 200,
		Body:       string(data),
	}

	return res, nil
}

func main() {
	lambda.Start(Handler)
}
