package bot

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/checkmateafrica/users/pkg/store"
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

	message := craftSignupMessage(user.UserID)

	if _, _, err := api.PostMessage(data.User, message); err != nil {
		log.Println(err)
		return
	}
}

func ShowSignupModal(p slack.InteractionCallback) {
	modal := craftSignUpModal(p)

	if _, err := api.OpenView(p.TriggerID, modal); err != nil {
		log.Println(err)
		return
	}
}

func SaveSignupData(p slack.InteractionCallback) {
	values := p.View.State.Values
	form := signUpform

	var (
		skillCategory   []string
		experienceLevel = values[form.experienceLevel.blockId][form.experienceLevel.actionId].SelectedOption
		gender          = values[form.gender.blockId][form.gender.actionId].SelectedOption
	)

	selectedCategories := values[form.skillCategory.blockId][form.skillCategory.actionId].SelectedOptions
	for _, item := range selectedCategories {
		skillCategory = append(skillCategory, item.Value)
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
		Name:            profile.DisplayName,
		Email:           profile.Email,
		SkillCategory:   skillCategory,
		ExperienceLevel: experienceLevel.Value,
		Gender:          gender.Value,
	}

	originMsgParams := strings.Split(p.View.PrivateMetadata, MetedataSeperator)
	api.DeleteMessage(originMsgParams[0], originMsgParams[1])

	store.SaveUserData(userData)
	SendSignupSuccessMessage(user.UserID, p)
}

func SendSignupSuccessMessage(userId string, p slack.InteractionCallback) {
	message := craftSignupSuccessMessage(userId, p)

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
