package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/checkmateafrica/accountability-bot/pkg/bot"
	"github.com/checkmateafrica/accountability-bot/pkg/store"
	"github.com/checkmateafrica/accountability-bot/pkg/utils"
	"github.com/checkmateafrica/accountability-bot/services"
	"github.com/slack-go/slack"
)

func handler(req utils.InvokeRequestPayload) error {
	payload := req.InteractionPayload

	interactionHandlers := map[slack.InteractionType]func(slack.InteractionCallback){
		slack.InteractionTypeViewSubmission: handleViewSubmission,
		slack.InteractionTypeBlockActions:   handleBlockAction,
	}

	if _, found := interactionHandlers[payload.Type]; !found {
		log.Println("unhandled interaction type")
		return nil
	}

	handler := interactionHandlers[payload.Type]
	handler(payload)

	return nil
}

func handleViewSubmission(p slack.InteractionCallback) {
	switch p.View.CallbackID {
	case utils.BlockIdSignupButton:
		bot.SaveBackgroundData(p, true)
	case utils.BlockIdSettingsButton:
		bot.SaveBackgroundData(p, false)
	default:
		log.Println("unhandled submission")
	}
}

func handleBlockAction(p slack.InteractionCallback) {
	for _, blockAction := range p.ActionCallback.BlockActions {
		switch blockAction.ActionID {
		case utils.ActionIdOpenModal:
			payload := p
			payload.BlockID = blockAction.BlockID
			bot.ShowBackgroundDataModal(payload)
		default:
			log.Println("unhandled interaction")
		}
	}
}

func main() {
	log.SetPrefix("ERROR: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	store.DB = services.NewDynaClient()
	lambda.Start(handler)
}
