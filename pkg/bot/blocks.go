package bot

import (
	"strconv"
	"time"

	"github.com/gosimple/slug"
	"github.com/slack-go/slack"
)

var spacerBlock = slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", "\n", false, false), nil, nil)

func craftSignupMessage(userId string) slack.MsgOption {
	textBlocksContent := []string{
		"Hola <@" + userId + ">! :tada:",
		"Welcome to our community, we are thrilled to have you here, let's get you started by signing you up for our weekly accountability partner pairings.",
		"Please click the button below to continue.",
	}

	var textBlocks []*slack.SectionBlock

	for _, blockText := range textBlocksContent {
		paragraph := slack.NewTextBlockObject("mrkdwn", blockText, false, false)
		block := slack.NewSectionBlock(paragraph, nil, nil)

		textBlocks = append(textBlocks, block)
	}

	buttonArgs := slack.ButtonBlockElement{
		Type:     "button",
		Text:     slack.NewTextBlockObject("plain_text", "Proceed to sign up", false, false),
		ActionID: ActionSignupButtonClick,
		Style:    "primary",
	}

	button := slack.ButtonBlockElement(buttonArgs)
	buttonBlock := slack.NewActionBlock(button.ActionID, button)

	return slack.MsgOptionBlocks(
		textBlocks[0],
		textBlocks[1],
		textBlocks[2],
		spacerBlock,
		buttonBlock,
	)
}

func craftSignUpModal(p slack.InteractionCallback) slack.ModalViewRequest {
	const infoText = "Please fill out the form below. We'll match you up with people who have similar skillset and experience as you."

	skillCategorySelect := craftSelectField(signUpform.skillCategory)
	expereienceLevelSelect := craftSelectField(signUpform.experienceLevel)
	genderSelect := craftSelectField(signUpform.gender)

	infoParagraph := slack.NewTextBlockObject("mrkdwn", infoText, false, false)
	infoBlock := slack.NewSectionBlock(infoParagraph, nil, nil)

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			spacerBlock,
			infoBlock,
			spacerBlock,
			skillCategorySelect,
			spacerBlock,
			expereienceLevelSelect,
			spacerBlock,
			genderSelect,
			spacerBlock,
		},
	}

	titleText := slack.NewTextBlockObject("plain_text", "Accountability Sign up", false, false)
	closeButtonText := slack.NewTextBlockObject("plain_text", "Cancel", false, false)
	submitButtonText := slack.NewTextBlockObject("plain_text", "Submit", false, false)

	var modalRequest = slack.ModalViewRequest{
		Type:            slack.ViewType("modal"),
		Title:           titleText,
		Close:           closeButtonText,
		Submit:          submitButtonText,
		Blocks:          blocks,
		CallbackID:      CallbackSignupModal,
		ClearOnClose:    true,
		PrivateMetadata: p.Channel.ID + MetedataSeperator + p.Message.Timestamp,
	}

	return modalRequest
}

func craftSelectField(f FormField) *slack.InputBlock {
	var optionBlocks []*slack.OptionBlockObject

	for _, s := range f.options {
		label := slack.NewTextBlockObject("plain_text", s, false, false)
		newOption := slack.NewOptionBlockObject(slug.Make(s), label, nil)

		optionBlocks = append(optionBlocks, newOption)
	}

	placeholderText := slack.NewTextBlockObject("plain_text", f.placeholder, true, false)
	labeltext := slack.NewTextBlockObject("plain_text", f.label, false, false)

	var selectType string
	if f.multi {
		selectType = "multi_static_select"
	} else {
		selectType = "static_select"
	}

	selectElement := slack.NewOptionsSelectBlockElement(selectType, placeholderText, f.actionId, optionBlocks...)
	return slack.NewInputBlock(f.blockId, labeltext, nil, selectElement)
}

func craftSignupSuccessMessage(userId string, p slack.InteractionCallback) slack.MsgOption {
	var (
		currentDate     = time.Now()
		daysUntilSunday = (0 + 7 - int(currentDate.Weekday())) % 7
		nextSunday      = currentDate.AddDate(0, 0, daysUntilSunday)
		day             = strconv.Itoa(nextSunday.Day())
		month           = nextSunday.Month().String()
		year            = strconv.Itoa(nextSunday.Year())
		daysTime        = strconv.Itoa(daysUntilSunday)
	)

	var nextRotation = "*Sunday, " + day + " " + month + ", " + year + "*"

	textBlocksContent := []string{
		"You're all set <@" + userId + ">!",
		"You've been added to our accountability partner rotations, you'll be paired with someone in your skill category every week. Next partner pairing will be done on " + nextRotation + " which is in *" + daysTime + " days time.*",
		"For now, you can introduce yourself in " + ChannelIntroductions + " or join one of the following channels to connect with others.",
		">" + ChannelDesign + "\n>" + ChannelEngineering + "\n>" + ChannelSecurityCompliance + "\n>" + ChannelContentManagement + "\n>" + ChannelDataAi,
	}

	var textBlocks []*slack.SectionBlock

	for _, blockText := range textBlocksContent {
		paragraph := slack.NewTextBlockObject("mrkdwn", blockText, false, false)
		block := slack.NewSectionBlock(paragraph, nil, nil)

		textBlocks = append(textBlocks, block)
	}

	return slack.MsgOptionBlocks(
		textBlocks[0],
		textBlocks[1],
		textBlocks[2],
		textBlocks[3],
	)
}

// func craftHomeScreenData() {

// }
