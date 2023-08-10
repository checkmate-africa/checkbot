package handlers

import (
	"fmt"

	"github.com/checkmateafrica/accountability-bot/pkg/bot"
	"github.com/checkmateafrica/accountability-bot/pkg/utils"
	"github.com/slack-go/slack"
)

func handleViewSubmission(p slack.InteractionCallback) {
	switch p.View.CallbackID {
	case utils.BlockIdSignupButton:
		bot.SaveBackgroundData(p, true)
	case utils.BlockIdSettingsButton:
		bot.SaveBackgroundData(p, false)
	default:
		fmt.Println("unhandled submission", p.View.CallbackID)
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
			fmt.Println("unhandled interaction")
		}
	}
}
