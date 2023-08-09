package blocks

import "github.com/gosimple/slug"

var skillCategoryOptions = []string{
	"3D Design",
	"Brand Design",
	"Product/UI/UX Design",
	"Motion Design",
	"Frontend Development",
	"Backend Development",
	"Mobile Development",
	"Cloud Engineering",
	"Data Science/AI/ML",
	"Cyber Security",
	"DevOps/SRE",
	"Writing/Technical Writing",
	"Product Management",
	"Developer Advocacy",
	"Art and Illustrations",
	"Webflow/Wordpress",
	"Legal/Tech Law",
}

var expereienceLevelOptions = []string{
	"Beginner",
	"Intermediate/Mid Level",
	"Senior",
	"Principal/Manegerial",
}

var genderOptions = []string{
	"Male",
	"Female",
	"Non Binary",
	"Prefer not to say",
}

const (
	labelSkillCategory   = "Skill Category"
	labelExperienceLevel = "Experience Level"
	labelGender          = "Gender"
)

var SignUpform = SignUpformType{
	SkillCategory: FormField{
		Placeholder: "Select one or more skills",
		BlockId:     "block" + slug.Make(labelSkillCategory),
		ActionId:    "action" + slug.Make(labelSkillCategory),
		Label:       labelSkillCategory,
		Options:     skillCategoryOptions,
		Multi:       true,
	},

	ExperienceLevel: FormField{
		Placeholder: "Select an option",
		BlockId:     "block" + slug.Make(labelExperienceLevel),
		ActionId:    "action" + slug.Make(labelExperienceLevel),
		Label:       labelExperienceLevel,
		Options:     expereienceLevelOptions,
		Multi:       false,
	},

	Gender: FormField{
		Placeholder: "Select an option",
		BlockId:     "block" + slug.Make(labelGender),
		ActionId:    "action" + slug.Make(labelGender),
		Label:       labelGender,
		Options:     genderOptions,
		Multi:       false,
	},
}
