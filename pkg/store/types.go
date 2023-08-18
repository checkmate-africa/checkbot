package store

type User struct {
	Email           string
	Name            string
	SkillCategories []string
	ExperienceLevel string
	Gender          string
	SlackId         string
}

type Pairs [][]User

type PairedUser struct {
	Email   string
	SlackId string
	Partner User
}
