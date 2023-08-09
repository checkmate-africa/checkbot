package bot

type EventItem struct {
	Ts string `json:"ts"`
}

type Event struct {
	Reaction string    `json:"reaction"`
	User     string    `json:"user"`
	Item     EventItem `json:"item"`
}

type EventData struct {
	Event `json:"event"`
}
