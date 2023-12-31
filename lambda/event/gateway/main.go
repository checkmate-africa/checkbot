package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	lambdaruntime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/checkmateafrica/accountability-bot/internal/bot"
	"github.com/checkmateafrica/accountability-bot/internal/utils"
	"github.com/checkmateafrica/accountability-bot/services"
	"github.com/slack-go/slack/slackevents"
)

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	body := req.Body
	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())

	if err != nil {
		log.Println(err)
		return utils.ApiResponse(http.StatusBadRequest, nil)
	}

	if eventsAPIEvent.Type == slackevents.URLVerification {
		res := bot.VerifyUrl(body)
		return utils.ApiResponse(http.StatusOK, &res.Challenge)
	}

	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		svc := services.NewLambdaService()

		invokePayload, err := json.Marshal(utils.InvokeRequestPayload{
			Body: body,
		})

		if err != nil {
			log.Println(err)
			return utils.ApiResponse(http.StatusBadRequest, nil)
		}

		input := &lambda.InvokeInput{
			FunctionName:   aws.String(utils.LambdaEventTaskFunction),
			InvocationType: aws.String("Event"),
			Payload:        invokePayload,
		}

		if _, err = svc.Invoke(input); err != nil {
			log.Println(err)
			return utils.ApiResponse(http.StatusInternalServerError, nil)
		}
	}

	return utils.ApiResponse(http.StatusOK, nil)
}

func main() {
	log.SetPrefix("ERROR: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	lambdaruntime.Start(handler)
}
