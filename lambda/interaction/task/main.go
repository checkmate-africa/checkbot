package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/checkmateafrica/accountability-bot/internal/bot"
	"github.com/checkmateafrica/accountability-bot/internal/store"
	"github.com/checkmateafrica/accountability-bot/internal/utils"
	"github.com/checkmateafrica/accountability-bot/services"
	"github.com/slack-go/slack"
)

var api = slack.New(os.Getenv(utils.EnvSlackToken))

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
		case utils.ActionIdToggleRetention:
			handleRetnetionToggle(p)
		default:
			log.Println("unhandled interaction")
		}
	}
}

func handleRetnetionToggle(p slack.InteractionCallback) {
	params := slack.GetUserProfileParameters{
		UserID:        p.User.ID,
		IncludeLabels: false,
	}

	profile, err := api.GetUserProfile(&params)
	if err != nil {
		log.Println(err)
		return
	}

	user, err := store.GetUser(profile.Email)
	if err != nil {
		log.Println(err)
		return
	}

	newRetention := !user.RetainPartner
	err = store.UpdateUserRentention(profile.Email, newRetention)
	if err != nil {
		log.Println(err)
		return
	}

	bot.PublishAppHome(p.User.ID, false)
}

func main() {
	log.SetPrefix("ERROR: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	store.DB = services.NewDynaClient()
	lambda.Start(handler)
}
