package store

type User struct {
	Email           string
	Name            string
	SkillCategories []string
	ExperienceLevel string
	Gender          string
	SlackId         string
	RetainPartner   bool
}

type Pairs [][]User

type PairedUser struct {
	Email   string
	SlackId string
	Partner User
}
