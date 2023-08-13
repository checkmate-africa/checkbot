package utils

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/slack-go/slack"
)

type InvokeRequestPayload struct {
	Body               string                    `json:"body"`
	InteractionPayload slack.InteractionCallback `json:"interactionPayload"`
}

func ApiResponse(status int, body interface{}) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{Headers: map[string]string{"Content-Type": "application/json"}}
	resp.StatusCode = status

	stringBody, _ := json.Marshal(body)
	resp.Body = string(stringBody)

	return &resp, nil
}
