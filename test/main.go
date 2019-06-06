package main

import (
	"github.com/k0kubun/pp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	tableName = "meocon_contacts"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String("ap-southeast-1")},
	}))

	db := dynamodb.New(sess)

	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("meocon_contacts"),
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String("ngominhluanbox@gmail.com"),
			},
		},
	})
	if err != nil {
		panic(err)
	}

	pp.Println(result.Item)
}
