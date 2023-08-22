package bot

type ReactionEventItem struct {
	Ts string `json:"ts"`
}

type SlackReactionEvent struct {
	Reaction string            `json:"reaction"`
	User     string            `json:"user"`
	Item     ReactionEventItem `json:"item"`
}

type SlackReactionEventData struct {
	SlackReactionEvent `json:"event"`
}

type JoinEventUser struct {
	Id string `json:"id"`
}

type SlackJoinEvent struct {
	User JoinEventUser `json:"user"`
}

type SlackJoinEventData struct {
	SlackJoinEvent `json:"event"`
}
