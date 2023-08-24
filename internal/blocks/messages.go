package blocks

import (
	"github.com/checkmateafrica/accountability-bot/internal/store"
	"github.com/checkmateafrica/accountability-bot/internal/utils"
	"github.com/slack-go/slack"
)

func SignupInviteMessage(userId string) slack.MsgOption {
	textContent :=
		"Hola <@" + userId + ">! :tada: \n\nWelcome to checkmate, we are thrilled to have you here in our community. Our primary goal here is to stay productive, skill up, connect and collaborate with peers. \n\nLet's get you signed up for accountability pairings."

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

func SignupSuccessMessage(userId string) slack.MsgOption {
	day, month, year, daysTime := utils.GetNextDayOfWeek(0)
	var nextRotation = "*Sunday, " + day + " " + month + ", " + year + "*"

	textContent :=
		"You're all set! <@" + userId + "> \n\nYou've been added to our accountability partner rotations, you'll be paired with someone in your skill category every week. Next pair shuffle will be done on " + nextRotation + " which is in *" + daysTime + " days time.* \n\nFor now, you can introduce yourself in " + utils.ChannelIntroductions + " or join one of the following channels to connect with others. \n\n>" + utils.ChannelDesign + "\n>" + utils.ChannelEngineering + "\n>" + utils.ChannelSecurityCompliance + "\n>" + utils.ChannelContentManagement + "\n>" + utils.ChannelDataAi

	paragraphs := slack.NewTextBlockObject("mrkdwn", textContent, false, false)
	textBlock := slack.NewSectionBlock(paragraphs, nil, nil)

	return slack.MsgOptionBlocks(
		textBlock,
	)
}

func PairShuffleAnnouncementMessage(pairs store.Pairs) slack.MsgOption {
	introText := "Unveiling This Week's Dynamic Duos! :tada: \n\nHi checkers! Your favourite bot is here again with the weekly shuffle and I am pleased to announce that I have matched y'all with new accountability partners according to your skill categories. \n\n Say goodbye to old buddies and prepare to collaborate, learn, share knowledge and stay productive with someone new. \n\n \n"
	closingText := "\n\n \nRemember to be awesome, respectful and not violate our community guidelines because we'll kick you out if you do (hehe not kidding), Have an amazing week ahead! :rocket:"

	introParagraph := slack.NewTextBlockObject("mrkdwn", introText, false, false)
	introBlock := slack.NewSectionBlock(introParagraph, nil, nil)

	var pairListText string

	for _, pair := range pairs {
		var text string

		if len(pair) > 1 {
			text = ":handshake: <@" + pair[0].SlackId + "> & <@" + pair[1].SlackId + "> \n"
		} else {
			text = "\n\n:ninja: <@" + pair[0].SlackId + "> is the lone wolf for the week. \n"
		}

		pairListText = pairListText + text
	}

	pairListParagraph := slack.NewTextBlockObject("mrkdwn", pairListText, false, false)
	pairListBlock := slack.NewSectionBlock(pairListParagraph, nil, nil)

	closingParagraph := slack.NewTextBlockObject("mrkdwn", closingText, false, false)
	closingBlock := slack.NewSectionBlock(closingParagraph, nil, nil)

	return slack.MsgOptionBlocks(
		introBlock,
		pairListBlock,
		closingBlock,
	)
}

func PairNotificationMessage(user store.PairedUser) slack.MsgOption {
	text := "Hi there :wave: \n\n Your new accountability partner for this week is <@" + user.Partner.SlackId + ">. \n\n You can navigate to the home tab to view more details about them. Remember to be respectful and adhere to our community guidelines. Have a productive week ahead!"

	paragraph := slack.NewTextBlockObject("mrkdwn", text, false, false)
	block := slack.NewSectionBlock(paragraph, nil, nil)

	return slack.MsgOptionBlocks(block)
}
