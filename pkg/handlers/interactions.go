package handlers

import (
	"fmt"

	"github.com/checkmateafrica/users/pkg/bot"
	"github.com/slack-go/slack"
)

func handleViewSubmission(p slack.InteractionCallback) {
	switch p.View.CallbackID {
	case bot.CallbackSignupModal:
		bot.SaveSignupData(p)
	default:
		fmt.Println("unhandled submission", p.View.CallbackID)
	}
}

func handleBlockAction(p slack.InteractionCallback) {
	for _, blockAction := range p.ActionCallback.BlockActions {
		switch blockAction.ActionID {
		case bot.ActionSignupButtonClick:
			bot.ShowSignupModal(p)
		default:
			fmt.Println("unhandled interaction")
		}
	}
}
