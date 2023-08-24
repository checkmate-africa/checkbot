package services

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/checkmateafrica/accountability-bot/internal/utils"
)

func NewDynaClient() *dynamodb.DynamoDB {
	config := &aws.Config{
		Region: aws.String(os.Getenv(utils.EnvAwsRegion)),
	}

	if utils.IsLocalEnv() {
		config.Endpoint = aws.String("http://host.docker.internal:8000")
	}

	awsSess, err := session.NewSession(config)

	if err != nil {
		log.Println(err)
		return nil
	}

	return dynamodb.New(awsSess)
}
