package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/checkmateafrica/accountability-bot/pkg/bot"
	"github.com/checkmateafrica/accountability-bot/pkg/store"
	"github.com/checkmateafrica/accountability-bot/pkg/utils"
	"github.com/checkmateafrica/accountability-bot/services"
	"github.com/slack-go/slack/slackevents"
)

func handler(req utils.InvokeRequestPayload) error {
	eventsAPIEvent, _ := slackevents.ParseEvent(json.RawMessage(req.Body), slackevents.OptionNoVerifyToken())
	innerEvent := eventsAPIEvent.InnerEvent.Data

	switch innerEvent.(type) {
	case *slackevents.TeamJoinEvent:
		var data bot.SlackJoinEventData

		if err := json.Unmarshal([]byte(req.Body), &data); err != nil {
			log.Println(err)
			return err
		}

		bot.PublishAppHome(data.User.Id)
		bot.InviteToSignup(data.User.Id)
	case *slackevents.ReactionAddedEvent:
		bot.DeleteMessageByReaction(req.Body)
	default:
		log.Println("unhandled event")
	}

	return nil
}

func main() {
	log.SetPrefix("ERROR: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	store.DB = services.NewDynaClient()
	lambda.Start(handler)
}
