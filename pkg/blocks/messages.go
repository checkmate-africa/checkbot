package blocks

import (
	"github.com/checkmateafrica/users/pkg/utils"
	"github.com/slack-go/slack"
)

func SignupInviteMessage(userId string) slack.MsgOption {
	textContent :=
		"Hola <@" + userId + ">! :tada: \n\n Welcome to checkmate, we are thrilled to have you here in our community, whether you're a pro, a curious learner, or somewhere in between, you're in the right place. \n\n \n\n Let's get you signed up for accountability pairings."

	paragraphs := slack.NewTextBlockObject("mrkdwn", textContent, false, false)
	textBlock := slack.NewSectionBlock(paragraphs, nil, nil)

	buttonArgs := slack.ButtonBlockElement{
		Type:     "button",
		Style:    "primary",
		Text:     slack.NewTextBlockObject("plain_text", "Proceed to sign up", false, false),
		ActionID: utils.ActionIdOpenModal,
	}

	button := slack.ButtonBlockElement(buttonArgs)
	buttonBlock := slack.NewActionBlock(utils.BlockIdSignupButton, button)
	spacer := slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", "\n", false, false), nil, nil)

	return slack.MsgOptionBlocks(
		textBlock,
		spacer,
		buttonBlock,
	)
}

func SignupSuccessMessage(userId string, p slack.InteractionCallback) slack.MsgOption {
	day, month, year, daysTime := utils.GetNextDayOfWeek(0)
	var nextRotation = "*Sunday, " + day + " " + month + ", " + year + "*"

	textContent :=
		"You're all set! <@" + userId + "> \n\n You've been added to our accountability partner rotations, you'll be paired with someone in your skill category every week. Next pair shuffle will be done on " + nextRotation + " which is in *" + daysTime + " days time.* \n\n For now, you can introduce yourself in " + utils.ChannelIntroductions + " or join one of the following channels to connect with others. \n\n>" + utils.ChannelDesign + "\n>" + utils.ChannelEngineering + "\n>" + utils.ChannelSecurityCompliance + "\n>" + utils.ChannelContentManagement + "\n>" + utils.ChannelDataAi

	paragraphs := slack.NewTextBlockObject("mrkdwn", textContent, false, false)
	textBlock := slack.NewSectionBlock(paragraphs, nil, nil)

	return slack.MsgOptionBlocks(
		textBlock,
	)
}
