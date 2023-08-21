package services

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/checkmateafrica/accountability-bot/pkg/utils"
)

func NewLambdaService() *lambda.Lambda {
	config := &aws.Config{
		Region: aws.String(os.Getenv(utils.EnvAwsRegion)),
	}

	if os.Getenv(utils.EnvSamLocal) == "true" {
		config.Endpoint = aws.String("http://host.docker.internal:3001")
	}

	awsSess, err := session.NewSession(config)

	if err != nil {
		log.Println(err)
		return nil
	}

	return lambda.New(awsSess)
}
