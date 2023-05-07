package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"os"
)

func PutItem(item map[string]*dynamodb.AttributeValue) (*dynamodb.PutItemOutput, error) {
	client := Client()
	_, err := client.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("TABLE_NAME")),
		Item:      item,
	})

	if err != nil {
		return nil, err
	}

	return nil, nil
}