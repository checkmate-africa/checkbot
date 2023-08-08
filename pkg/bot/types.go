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

type FormField struct {
	placeholder string
	blockId     string
	actionId    string
	label       string
	options     []string
	multi       bool
}

type SignUpform struct {
	skillCategory   FormField
	experienceLevel FormField
	gender          FormField
}
