package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/k0kubun/pp"
)

var (
	tableName = "meocon_contacts"
)

type contacts struct {
	Name    string
	Email   string
	Message string
}

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            aws.Config{Region: aws.String("ap-southeast-1")},
		SharedConfigState: session.SharedConfigEnable,
	}))

	db := dynamodb.New(sess)

	result, err := db.Scan(&dynamodb.ScanInput{
		TableName: aws.String("meocon_contacts"),
		Limit:     aws.Int64(2),
		Segment:   aws.Int64(2),
	})
	if err != nil {
		panic(err)
	}

	c := []contacts{}
	for _, item := range result.Items {
		contact := contacts{
			Name:    *item["name"].S,
			Email:   *item["email"].S,
			Message: *item["message"].S,
		}
		c = append(c, contact)
	}

	pp.Println(c)
}
