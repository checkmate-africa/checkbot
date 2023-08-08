package bot

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
	"Technical Writing",
	"Product Management",
	"Developer Advocacy",
	"Illustrator/Art",
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

var signUpform = SignUpform{
	skillCategory: FormField{
		placeholder: "Select one or more skills",
		blockId:     "block" + slug.Make(labelSkillCategory),
		actionId:    "action" + slug.Make(labelSkillCategory),
		label:       labelSkillCategory,
		options:     skillCategoryOptions,
		multi:       true,
	},

	experienceLevel: FormField{
		placeholder: "Select an option",
		blockId:     "block" + slug.Make(labelExperienceLevel),
		actionId:    "action" + slug.Make(labelExperienceLevel),
		label:       labelExperienceLevel,
		options:     expereienceLevelOptions,
		multi:       false,
	},

	gender: FormField{
		placeholder: "Select an option",
		blockId:     "block" + slug.Make(labelGender),
		actionId:    "action" + slug.Make(labelGender),
		label:       labelGender,
		options:     genderOptions,
		multi:       false,
	},
}
