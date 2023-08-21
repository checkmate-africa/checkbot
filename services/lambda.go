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
	awsSess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String("http://host.docker.internal:3001"),
		Region:   aws.String(os.Getenv(utils.AwsRegion)),
	})

	if err != nil {
		log.Println(err)
		return nil
	}

	return lambda.New(awsSess)
}
