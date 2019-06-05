package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	uuid "github.com/satori/go.uuid"
)

var (
	tableName = "meocon-development"
)

func init() {
	tableName = os.Getenv("DYNAMODB_TABLE")
}

// Response alias Response
type Response events.APIGatewayProxyResponse

// Request alias Request
type Request events.APIGatewayProxyRequest

type userInfo struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type item struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
	ID      string `json:"id"`
}

func errorResponse(e error) Response {
	return Response{
		StatusCode: 400,
		Body:       e.Error(),
	}
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request Request) (Response, error) {
	u := userInfo{}
	if err := json.Unmarshal([]byte(request.Body), &u); err != nil {
		return Response{
			StatusCode: 400,
			Body:       err.Error(),
		}, nil
	}

	if (u.Name == "") || (u.Email == "") || (u.Message == "") {
		return Response{
			StatusCode: 400,
			Body:       fmt.Sprintf("{\"status\": \"failure\""),
		}, nil
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)

	id := uuid.NewV4()
	i := item{
		ID:      id.String(),
		Name:    u.Name,
		Email:   u.Email,
		Message: u.Message,
	}

	av, err := dynamodbattribute.MarshalMap(i)
	if err != nil {
		return errorResponse(err), nil
	}

	_, err = svc.PutItem(&dynamodb.PutItemInput{
		Item:      av,
		TableName: &tableName,
	})
	if err != nil {
		return errorResponse(err), nil
	}

	return Response{
		StatusCode: 200,
		Body:       fmt.Sprintf("{\"status\": \"ok\""),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
