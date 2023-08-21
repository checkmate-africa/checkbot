package services

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/checkmateafrica/accountability-bot/pkg/utils"
)

func NewDynaClient() *dynamodb.DynamoDB {
	awsSess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String("http://host.docker.internal:8000"),
		Region:   aws.String(os.Getenv(utils.AwsRegion)),
	})

	if err != nil {
		log.Println(err)
		return nil
	}

	return dynamodb.New(awsSess)
}
