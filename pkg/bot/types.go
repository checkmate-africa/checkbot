package bot

type EventItem struct {
	Ts string `json:"ts"`
}

type SlackEvent struct {
	Reaction string    `json:"reaction"`
	User     string    `json:"user"`
	Item     EventItem `json:"item"`
}

type SlackEventData struct {
	SlackEvent `json:"event"`
}
