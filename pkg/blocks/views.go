package blocks

import (
	"github.com/checkmateafrica/users/pkg/store"
	"github.com/checkmateafrica/users/pkg/utils"
	"github.com/gosimple/slug"
	"github.com/slack-go/slack"
)

func BackgroundDataModal(p slack.InteractionCallback) slack.ModalViewRequest {
	const infoText = "Please fill out the form below. We'll match you up with people who have similar skillset and experience as you."

	skillCategorySelect := selectField(SignUpform.SkillCategory)
	expereienceLevelSelect := selectField(SignUpform.ExperienceLevel)
	genderSelect := selectField(SignUpform.Gender)

	infoParagraph := slack.NewTextBlockObject("mrkdwn", infoText, false, false)
	infoBlock := slack.NewSectionBlock(infoParagraph, nil, nil)

	spacer := slack.NewSectionBlock(slack.NewTextBlockObject("mrkdwn", "\n", false, false), nil, nil)

	blocks := slack.Blocks{
		BlockSet: []slack.Block{
			spacer,
			infoBlock,
			spacer,
			skillCategorySelect,
			spacer,
			expereienceLevelSelect,
			spacer,
			genderSelect,
			spacer,
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
		CallbackID:      p.BlockID,
		ClearOnClose:    true,
		PrivateMetadata: p.Channel.ID + utils.MetedataSeperator + p.Message.Timestamp,
	}

	return modalRequest
}

func selectField(f FormField) *slack.InputBlock {
	var optionBlocks []*slack.OptionBlockObject

	for _, s := range f.Options {
		label := slack.NewTextBlockObject("plain_text", s, false, false)
		newOption := slack.NewOptionBlockObject(slug.Make(s), label, nil)

		optionBlocks = append(optionBlocks, newOption)
	}

	placeholderText := slack.NewTextBlockObject("plain_text", f.Placeholder, true, false)
	labeltext := slack.NewTextBlockObject("plain_text", f.Label, false, false)

	var selectType string
	if f.Multi {
		selectType = "multi_static_select"
	} else {
		selectType = "static_select"
	}

	selectElement := slack.NewOptionsSelectBlockElement(selectType, placeholderText, f.ActionId, optionBlocks...)
	return slack.NewInputBlock(f.BlockId, labeltext, nil, selectElement)
}

func AppHomeContent(partner *store.User) slack.HomeTabViewRequest {
	var blocks slack.Blocks

	headerCtaButtonText := slack.NewTextBlockObject("plain_text", "Settings", false, false)
	headerCtaElement := slack.NewButtonBlockElement(utils.ActionIdOpenModal, "", headerCtaButtonText)
	headerCta := slack.NewAccessory(headerCtaElement)

	_, _, _, daysTime := utils.GetNextDayOfWeek(0)

	headerText := slack.NewTextBlockObject("mrkdwn", "*Accountability Partner* \nNext rotation in _"+daysTime+" days time_.", false, false)
	headerArgs := slack.SectionBlockOptionBlockID(utils.BlockIdSettingsButton)

	header := slack.NewSectionBlock(headerText, nil, headerCta, headerArgs)
	divider := slack.NewDividerBlock()

	if partner == nil {
		blocks = slack.Blocks{
			BlockSet: []slack.Block{
				header,
			},
		}
	} else {
		profileBlocks := partnerProfileBlocks(partner)

		blocks = slack.Blocks{
			BlockSet: append([]slack.Block{
				header,
				divider,
			}, profileBlocks...),
		}
	}

	var viewRequest = slack.HomeTabViewRequest{
		Type:   slack.ViewType("home"),
		Blocks: blocks,
	}

	return viewRequest
}

func partnerProfileBlocks(partner *store.User) (blocks []slack.Block) {
	titleText := slack.NewTextBlockObject("mrkdwn", partner.Name, false, false)
	title := slack.NewHeaderBlock(titleText)

	return []slack.Block{title}
}
