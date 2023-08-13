package services

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func NewLambdaService() *lambda.Lambda {
	awsSess, err := session.NewSession(&aws.Config{
		Region:      aws.String("local"),
		Endpoint:    aws.String("http://host.docker.internal:3001"),
		Credentials: credentials.NewStaticCredentials("x", "x", "xxxxx"),
	})

	if err != nil {
		log.Println(err)
		return nil
	}

	return lambda.New(awsSess)
}
