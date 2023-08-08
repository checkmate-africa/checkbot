package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/checkmateafrica/users/pkg/bot"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func DefaultHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Kwame is running!"))
}

func EventsHandler(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	eventsAPIEvent, err := slackevents.ParseEvent(json.RawMessage(body), slackevents.OptionNoVerifyToken())

	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)

	if eventsAPIEvent.Type == slackevents.URLVerification {
		bot.VerifyUrl(w, body)
		return
	}

	if eventsAPIEvent.Type == slackevents.CallbackEvent {
		innerEvent := eventsAPIEvent.InnerEvent.Data

		switch innerEvent.(type) {
		// change to team joined before deploy
		case *slackevents.PinAddedEvent:
			go bot.InviteToSignup(body)
		case *slackevents.ReactionAddedEvent:
			go bot.DeleteMessageByReaction(body)
		default:
			fmt.Println("unhandled event")
			return
		}
	}

}

func InteractionsHandler(w http.ResponseWriter, req *http.Request) {
	var payload slack.InteractionCallback

	if err := json.Unmarshal([]byte(req.FormValue("payload")), &payload); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	interactionHandlers := map[slack.InteractionType]func(slack.InteractionCallback){
		slack.InteractionTypeViewSubmission: handleViewSubmission,
		slack.InteractionTypeBlockActions:   handleBlockAction,
	}

	if _, found := interactionHandlers[payload.Type]; !found {
		fmt.Println("unhandled interactionType:", payload.Type)
		return
	}

	handler := interactionHandlers[payload.Type]
	go handler(payload)
}
