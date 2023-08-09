package blocks

import (
	"github.com/gosimple/slug"
	"github.com/slack-go/slack"
)

var skillCategoryOptions = []string{
	"3D Design",
	"Brand Design",
	"Product/UI/UX Design",
	"Motion Design",
	"Frontend Development",
	"Backend Development",
	"Mobile Development",
	"Cloud Engineering",
	"Data Science/AI/ML",
	"Cyber Security",
	"DevOps/SRE",
	"Writing/Technical Writing",
	"Product Management",
	"Developer Advocacy",
	"Art and Illustrations",
	"Webflow/Wordpress",
	"Legal/Tech Law",
}

var expereienceLevelOptions = []string{
	"Beginner",
	"Intermediate/Mid Level",
	"Senior",
	"Principal/Manegerial",
}

var genderOptions = []string{
	"Male",
	"Female",
	"Non Binary",
	"Prefer not to say",
}

const (
	labelSkillCategory   = "Skill Category"
	labelExperienceLevel = "Experience Level"
	labelGender          = "Gender"
)

var SignUpform = SignUpformType{
	SkillCategory: FormField{

		Placeholder: "Select one or more skills",
		Hint:        "You can select more than one",
		BlockId:     "block" + slug.Make(labelSkillCategory),
		ActionId:    "action" + slug.Make(labelSkillCategory),
		Label:       labelSkillCategory,
		Options:     skillCategoryOptions,
		Multi:       true,
	},

	ExperienceLevel: FormField{
		Placeholder: "Select an option",
		Hint:        "Don't worry, pairing won't be 'strictly' limited to your experience level",
		BlockId:     "block" + slug.Make(labelExperienceLevel),
		ActionId:    "action" + slug.Make(labelExperienceLevel),
		Label:       labelExperienceLevel,
		Options:     expereienceLevelOptions,
		Multi:       false,
	},

	Gender: FormField{
		Placeholder: "Select an option",
		Hint:        "We use this information to make our community a safe space for all",
		BlockId:     "block" + slug.Make(labelGender),
		ActionId:    "action" + slug.Make(labelGender),
		Label:       labelGender,
		Options:     genderOptions,
		Multi:       false,
	},
}

func selectField(f FormField, initialOption *string, initialOptions *[]string) *slack.InputBlock {
	var optionBlocks []*slack.OptionBlockObject

	for _, s := range f.Options {
		label := slack.NewTextBlockObject("plain_text", s, false, false)
		newOption := slack.NewOptionBlockObject(s, label, nil)
		optionBlocks = append(optionBlocks, newOption)
	}

	placeholderText := slack.NewTextBlockObject("plain_text", f.Placeholder, true, false)
	labeltext := slack.NewTextBlockObject("plain_text", f.Label, false, false)
	hintText := slack.NewTextBlockObject("plain_text", f.Hint, false, false)

	var selectElement slack.BlockElement

	if f.Multi {
		multiSelect := slack.MultiSelectBlockElement{
			Placeholder: placeholderText,
			ActionID:    f.ActionId,
			Options:     optionBlocks,
			Type:        "multi_static_select",
		}

		if initialOptions != nil {
			var selectedOptionBlocks []*slack.OptionBlockObject

			for _, option := range *initialOptions {
				for _, block := range optionBlocks {
					if block.Text.Text == option {
						selectedOptionBlocks = append(selectedOptionBlocks, block)
						break
					}
				}
			}

			multiSelect.InitialOptions = selectedOptionBlocks
		}

		selectElement = multiSelect
	} else {
		singleSelect := slack.SelectBlockElement{
			Placeholder: placeholderText,
			ActionID:    f.ActionId,
			Options:     optionBlocks,
			Type:        "static_select",
		}

		if initialOption != nil {
			label := slack.NewTextBlockObject("plain_text", *initialOption, false, false)
			optionBlock := slack.NewOptionBlockObject(*initialOption, label, nil)
			singleSelect.InitialOption = optionBlock
		}

		selectElement = singleSelect
	}

	return slack.NewInputBlock(f.BlockId, labeltext, hintText, selectElement)
}
