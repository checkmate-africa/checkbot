package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/checkmateafrica/accountability-bot/internal/bot"
	"github.com/checkmateafrica/accountability-bot/internal/store"
	"github.com/checkmateafrica/accountability-bot/internal/utils"
	"github.com/checkmateafrica/accountability-bot/services"
	"github.com/slack-go/slack/slackevents"
)

func handler(req utils.InvokeRequestPayload) error {
	eventsAPIEvent, _ := slackevents.ParseEvent(json.RawMessage(req.Body), slackevents.OptionNoVerifyToken())
	innerEvent := eventsAPIEvent.InnerEvent.Data

	switch innerEvent.(type) {
	case *slackevents.TeamJoinEvent:
		var data bot.TeamJoinEventData

		if err := json.Unmarshal([]byte(req.Body), &data); err != nil {
			log.Println(err)
			return err
		}

		bot.InviteToSignup(data.User.ID)
	case *slackevents.MemberJoinedChannelEvent:
		var data bot.ChannelJoinEventData

		if err := json.Unmarshal([]byte(req.Body), &data); err != nil {
			log.Println(err)
			return err
		}

		if data.Channel == utils.ChannelIdManualSignup {
			bot.InviteToSignup(data.User)
		}
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
