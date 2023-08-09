package blocks

import (
	"strings"

	"github.com/checkmateafrica/users/pkg/store"
	"github.com/checkmateafrica/users/pkg/utils"
	"github.com/slack-go/slack"
)

func BackgroundDataModal(p slack.InteractionCallback, user *store.User) slack.ModalViewRequest {
	const (
		desc1 = "Please fill out the form below. We'll match you up with people who have similar skillset and experience as you."
		desc2 = "Update your accountability profile settings. Changes will be used in subsequest partner shuffles."
	)

	var (
		selectedCategories *[]string
		selectedLevel      *string
		selectedGender     *string
	)

	if user != nil {
		selectedCategories = &user.SkillCategories
		selectedLevel = &user.ExperienceLevel
		selectedGender = &user.Gender
	}

	skillCategorySelect := selectField(SignUpform.SkillCategory, nil, selectedCategories)
	expereienceLevelSelect := selectField(SignUpform.ExperienceLevel, selectedLevel, nil)
	genderSelect := selectField(SignUpform.Gender, selectedGender, nil)

	spacer := slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", "\n", false, false), nil, nil)

	modalDescriptions := map[string]string{
		utils.BlockIdSignupButton:   desc1,
		utils.BlockIdSettingsButton: desc2,
	}

	modalDescText := slack.NewTextBlockObject("mrkdwn", modalDescriptions[p.BlockID], false, false)
	modalDescBlock := slack.NewSectionBlock(modalDescText, nil, nil)

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			spacer,
			modalDescBlock,
			spacer,
			skillCategorySelect,
			spacer,
			expereienceLevelSelect,
			spacer,
			genderSelect,
			spacer,
		},
	}

	modalTitles := map[string]string{
		utils.BlockIdSignupButton:   "Accountability Sign up",
		utils.BlockIdSettingsButton: "Accountability Profile",
	}

	titleText := slack.NewTextBlockObject("plain_text", modalTitles[p.BlockID], false, false)
	closeButtonText := slack.NewTextBlockObject("plain_text", "Cancel", false, false)
	submitButtonText := slack.NewTextBlockObject("plain_text", "Submit", false, false)

	var modalRequest = slack.ModalViewRequest{
		Type:            slack.ViewType("modal"),
		Title:           titleText,
		Close:           closeButtonText,
		Submit:          submitButtonText,
		Blocks:          blocks,
		CallbackID:      p.BlockID,
		ClearOnClose:    true,
		PrivateMetadata: p.Channel.ID + utils.MetedataSeperator + p.Message.Timestamp,
	}

	return modalRequest
}

func AppHomeContent(partner *store.User) slack.HomeTabViewRequest {
	var blocks slack.Blocks

	headerCtaButtonText := slack.NewTextBlockObject("plain_text", "Settings", false, false)
	headerCtaElement := slack.NewButtonBlockElement(utils.ActionIdOpenModal, "", headerCtaButtonText)
	headerCta := slack.NewAccessory(headerCtaElement)

	headerArgs := slack.SectionBlockOptionBlockID(utils.BlockIdSettingsButton)
	divider := slack.NewDividerBlock()

	if partner == nil {
		headerText := slack.NewTextBlockObject("mrkdwn", "*Accountability* \nNo partner until next week.", false, false)
		header := slack.NewSectionBlock(headerText, nil, headerCta, headerArgs)

		blocks = slack.Blocks{
			BlockSet: []slack.Block{
				header,
			},
		}
	} else {
		headerText := slack.NewTextBlockObject("mrkdwn", "*Accountability* \nYour current partner.", false, false)
		header := slack.NewSectionBlock(headerText, nil, headerCta, headerArgs)

		profileBlocks := partnerProfileBlocks(partner)
		defaultBlocks := []slack.Block{
			header,
			divider,
		}

		blocks = slack.Blocks{
			BlockSet: append(defaultBlocks, profileBlocks...),
		}
	}

	var viewRequest = slack.HomeTabViewRequest{
		Type:   slack.ViewType("home"),
		Blocks: blocks,
	}

	return viewRequest
}

func partnerProfileBlocks(p *store.User) (blocks []slack.Block) {
	titleText := slack.NewTextBlockObject("plain_text", p.Name, false, false)
	title := slack.NewHeaderBlock(titleText)

	spacer := slack.NewTextBlockObject("mrkdwn", "\n", false, false)

	skillsfieldSlice := make([]*slack.TextBlockObject, 0)
	skillsText := slack.NewTextBlockObject("mrkdwn", "*Skill Categories:*\n_"+strings.Join(p.SkillCategories, ", ")+"_", false, false)
	skillsfieldSlice = append(append(skillsfieldSlice, skillsText), spacer)
	skillsBlock := slack.NewSectionBlock(nil, skillsfieldSlice, nil)

	expfieldSlice := make([]*slack.TextBlockObject, 0)
	expText := slack.NewTextBlockObject("mrkdwn", "*Experience Level:*\n_"+p.ExperienceLevel+"_", false, false)
	expfieldSlice = append(append(expfieldSlice, expText), spacer)
	expBlock := slack.NewSectionBlock(nil, expfieldSlice, nil)

	button1Args := slack.ButtonBlockElement{
		Type: "button",
		URL:  "https://google.com",
		Text: slack.NewTextBlockObject("plain_text", "Report abuse", false, false),
	}

	button2Args := slack.ButtonBlockElement{
		Type:  "button",
		Style: "primary",
		URL:   "slack://user?team=" + utils.TeamID + "&id=" + p.SlackId,
		Text:  slack.NewTextBlockObject("plain_text", "Send a message", false, false),
	}

	button1 := slack.ButtonBlockElement(button1Args)
	button2 := slack.ButtonBlockElement(button2Args)
	actionsBlock := slack.NewActionBlock(utils.BlockIdSignupButton, button1, button2)

	return []slack.Block{
		title,
		skillsBlock,
		expBlock,
		actionsBlock,
	}
}
