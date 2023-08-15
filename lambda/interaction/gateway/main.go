package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	lambdaruntime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/checkmateafrica/accountability-bot/pkg/utils"
	"github.com/checkmateafrica/accountability-bot/services"
	"github.com/gorilla/schema"
	"github.com/slack-go/slack"
)

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var interactionPayload slack.InteractionCallback

	decoded, _ := url.ParseQuery(req.Body)
	schema.NewDecoder().Decode(interactionPayload, decoded)

	if err := json.Unmarshal([]byte(decoded["payload"][0]), &interactionPayload); err != nil {
		return utils.ApiResponse(http.StatusBadRequest, "")
	}

	invokePayload, err := json.Marshal(utils.InvokeRequestPayload{
		InteractionPayload: interactionPayload,
	})

	if err != nil {
		log.Println(err)
		return utils.ApiResponse(http.StatusBadRequest, nil)
	}

	input := &lambda.InvokeInput{
		FunctionName:   aws.String("InteractionTaskFunction"),
		InvocationType: aws.String("RequestResponse"),
		Payload:        invokePayload,
	}

	svc := services.NewLambdaService()
	go svc.Invoke(input)

	return utils.ApiResponse(http.StatusOK, "success")
}

func main() {
	log.SetPrefix("ERROR: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	lambdaruntime.Start(handler)
}
