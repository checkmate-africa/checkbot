package blocks

type FormField struct {
	Placeholder string
	BlockId     string
	ActionId    string
	Label       string
	Hint        string
	Options     []string
	Multi       bool
}

type SignUpformType struct {
	SkillCategory   FormField
	ExperienceLevel FormField
	Gender          FormField
}
