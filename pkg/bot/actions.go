package bot

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/checkmateafrica/users/pkg/blocks"
	"github.com/checkmateafrica/users/pkg/store"
	"github.com/checkmateafrica/users/pkg/utils"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

var api = slack.New("xoxb-1465680528901-5671214476405-XpVXLaeSEoqdKCKPjZEDOpMH")

func VerifyUrl(w http.ResponseWriter, body []byte) {
	var res *slackevents.ChallengeResponse

	if err := json.Unmarshal(body, &res); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}

	w.Write(body)
}

func InviteToSignup(body []byte) {
	var data EventData

	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(err)
		return
	}

	user := slack.GetUserProfileParameters{
		UserID:        data.User,
		IncludeLabels: false,
	}

	message := blocks.SignupInviteMessage(user.UserID)

	if _, _, err := api.PostMessage(data.User, message); err != nil {
		log.Println(err)
		return
	}
}

func ShowBackgroundDataModal(p slack.InteractionCallback) {
	var user *store.User = nil

	if p.BlockID == utils.BlockIdSettingsButton {
		params := slack.GetUserProfileParameters{
			UserID:        p.User.ID,
			IncludeLabels: false,
		}

		profile, err := api.GetUserProfile(&params)

		if err != nil {
			log.Println(err)
		} else {
			user = store.GetUser(profile.Email)
		}
	}

	modal := blocks.BackgroundDataModal(p, user)

	if _, err := api.OpenView(p.TriggerID, modal); err != nil {
		log.Println(err)
		return
	}
}

func SaveBackgroundData(p slack.InteractionCallback, successMessage bool) {
	values := p.View.State.Values
	form := blocks.SignUpform

	var (
		skills          []string
		experienceLevel = values[form.ExperienceLevel.BlockId][form.ExperienceLevel.ActionId].SelectedOption
		gender          = values[form.Gender.BlockId][form.Gender.ActionId].SelectedOption
	)

	selectedCategories := values[form.SkillCategory.BlockId][form.SkillCategory.ActionId].SelectedOptions
	for _, item := range selectedCategories {
		skills = append(skills, item.Value)
	}

	user := slack.GetUserProfileParameters{
		UserID:        p.User.ID,
		IncludeLabels: false,
	}

	profile, err := api.GetUserProfile(&user)

	if err != nil {
		log.Println(err)
		return
	}

	userData := store.User{
		Email:           profile.Email,
		Name:            profile.RealName,
		SkillCategories: skills,
		ExperienceLevel: experienceLevel.Value,
		Gender:          gender.Value,
		SlackId:         user.UserID,
	}

	if successMessage {
		originMsgParams := strings.Split(p.View.PrivateMetadata, utils.MetedataSeperator)
		api.DeleteMessage(originMsgParams[0], originMsgParams[1])
		SendSignupSuccessMessage(user.UserID, p)
	}

	store.SaveUserData(userData)
}

func SendSignupSuccessMessage(userId string, p slack.InteractionCallback) {
	message := blocks.SignupSuccessMessage(userId, p)

	if _, _, err := api.PostMessage(userId, message); err != nil {
		log.Println(err)
		return
	}
}

func DeleteMessageByReaction(body []byte) {
	var data EventData

	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(err)
		return
	}

	if data.Reaction == "x" {
		if _, _, err := api.DeleteMessage(data.User, data.Item.Ts); err != nil {
			log.Println(err)
			return
		}
	}
}

func PublishAppHome(body []byte) {
	var data EventData

	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(err)
		return
	}

	user := slack.GetUserProfileParameters{
		UserID:        data.User,
		IncludeLabels: false,
	}

	profile, err := api.GetUserProfile(&user)

	if err != nil {
		log.Println(err)
		return
	}

	partner := store.GetPartner(profile.Email)
	view := blocks.AppHomeContent(partner)

	if _, err = api.PublishView(user.UserID, view, ""); err != nil {
		log.Println(err)
		return
	}
}
