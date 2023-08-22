package bot

import "github.com/slack-go/slack/slackevents"

type ReactionAddedEventData struct {
	slackevents.ReactionAddedEvent `json:"event"`
}

type TeamJoinEventData struct {
	slackevents.TeamJoinEvent `json:"event"`
}

type ChannelJoinEventData struct {
	slackevents.MemberJoinedChannelEvent `json:"event"`
}
