package bot

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/checkmateafrica/accountability-bot/internal/blocks"
	"github.com/checkmateafrica/accountability-bot/internal/store"
	"github.com/checkmateafrica/accountability-bot/internal/utils"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

var api = slack.New(os.Getenv(utils.EnvSlackToken))

func VerifyUrl(body string) *slackevents.ChallengeResponse {
	var res *slackevents.ChallengeResponse

	if err := json.Unmarshal([]byte(body), &res); err != nil {
		log.Println(err)
	}

	return res
}

func InviteToSignup(userId string) {
	message := blocks.SignupInviteMessage(userId)

	publishAppHome(userId, true)

	if _, _, err := api.PostMessage(userId, message); err != nil {
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
			user, _ = store.GetUser(profile.Email)
		}
	}

	modal := blocks.BackgroundDataModal(p, user)

	if _, err := api.OpenView(p.TriggerID, modal); err != nil {
		log.Println(err)
		return
	}
}

func SaveBackgroundData(p slack.InteractionCallback, isNewUser bool) {
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

	store.SaveUserData(userData)

	if isNewUser {
		originMsgParams := strings.Split(p.View.PrivateMetadata, utils.MetedataSeperator)
		if _, _, err := api.DeleteMessage(originMsgParams[0], originMsgParams[1]); err != nil {
			log.Println(err)
		}

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			SendSignupSuccessMessage(user.UserID, p)
		}()

		go func() {
			defer wg.Done()
			publishAppHome(user.UserID, true)
		}()

		wg.Wait()
	}
}

func SendSignupSuccessMessage(userId string, p slack.InteractionCallback) {
	message := blocks.SignupSuccessMessage(userId)

	if _, _, err := api.PostMessage(userId, message); err != nil {
		log.Println(err)
		return
	}
}

func DeleteMessageByReaction(body string) {
	var data ReactionAddedEventData

	if err := json.Unmarshal([]byte(body), &data); err != nil {
		log.Println(err)
		return
	}

	if data.Reaction == "x" {
		if _, _, err := api.DeleteMessage(data.User, data.Item.Timestamp); err != nil {
			log.Println(err)
			return
		}
	}
}

func SendNewPairShuffleAnnouncement(pairs store.Pairs) error {
	message := blocks.PairShuffleAnnouncementMessage(pairs)

	if _, _, err := api.PostMessage(utils.ChannelIdAnnouncements, message, slack.MsgOptionAsUser(slack.DEFAULT_MESSAGE_ASUSER)); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func SendPairShuffleNotification(pairedUsers []store.PairedUser) {
	var wg sync.WaitGroup

	for _, user := range pairedUsers {
		wg.Add(1)

		go func(u store.PairedUser) {
			defer wg.Done()
			message := blocks.PairNotificationMessage(u)

			publishAppHome(u.SlackId, false)

			if _, _, err := api.PostMessage(u.SlackId, message, slack.MsgOptionAsUser(slack.DEFAULT_MESSAGE_ASUSER)); err != nil {
				log.Println(err)
			}
		}(user)
	}

	wg.Wait()
}

func publishAppHome(userId string, isNewUser bool) {
	params := slack.GetUserProfileParameters{
		UserID:        userId,
		IncludeLabels: false,
	}

	profile, err := api.GetUserProfile(&params)

	if err != nil {
		log.Println(err)
		return
	}

	var user *store.User
	var partner *store.User

	if !isNewUser {
		user, _ = store.GetUser(profile.Email)
		partner, _ = store.GetPartner(profile.Email)
	}

	view := blocks.AppHomeContent(partner, user)

	if _, err = api.PublishView(userId, view, ""); err != nil {
		log.Println(err)
		return
	}
}
