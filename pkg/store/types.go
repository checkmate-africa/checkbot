package store

type User struct {
	Email           string
	Name            string
	SkillCategories []string
	ExperienceLevel string
	Gender          string
	SlackId         string
}

type Pair struct {
	Email   string
	Partner User
}
