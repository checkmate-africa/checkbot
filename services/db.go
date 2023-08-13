package services

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewDynaClient() *dynamodb.DynamoDB {
	awsSess, err := session.NewSession(&aws.Config{
		Region:      aws.String("local"),
		Endpoint:    aws.String("http://host.docker.internal:8000"),
		Credentials: credentials.NewStaticCredentials("x", "x", "xxxxx"),
	})

	if err != nil {
		log.Println(err)
		return nil
	}

	return dynamodb.New(awsSess)
}
