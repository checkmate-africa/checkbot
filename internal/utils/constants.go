package utils

const (
	MetedataSeperator             = "__"
	EnvAwsRegion                  = "AwsRegion"
	EnvSlackToken                 = "SlackBotToken"
	EnvSamLocal                   = "AWS_SAM_LOCAL"
	ActionIdOpenModal             = "actionId-openModal"
	BlockIdSignupButton           = "blockId-signupBtn"
	BlockIdSettingsButton         = "blockId-settingsBtn"
	TeamID                        = "T01DPL0FJSH"
	ChannelIdAnnouncements        = "C01DC0UHZST"
	ChannelIdManualSignup         = "C05P2T77QSY"
	ChannelIntroductions          = "<#C01EGJ730DN>"
	ChannelDesign                 = "<#C05LPDMJVV0>"
	ChannelEngineering            = "<#C05M2307FA5>"
	ChannelSecurityCompliance     = "<#C05L8T6EHPH>"
	ChannelContentManagement      = "<#C05LGQ591SS>"
	ChannelDataAi                 = "<#C05L8TCRN31>"
	LambdaEventTaskFunction       = "checkmate-EventTaskFunction"
	LambdaInteractionTaskFunction = "checkmate-InteractionTaskFunction"
)

var SkillDomains = map[string][]string{
	"design": {
		"3D Design",
		"Brand Design",
		"Product/UI/UX Design",
		"Motion Design",
		"Art and Illustrations",
	},

	"engineering": {
		"Frontend Development",
		"Backend Development",
		"Mobile Development",
		"Cloud Engineering",
		"Blockchain Engineering",
		"DevOps/SRE",
	},

	"data-science-and-ai": {
		"Data Science/AI/ML",
	},

	"content-and-management": {
		"Developer Advocacy",
		"Writing/Technical Writing",
		"Webflow/Wordpress",
		"Product Management",
		"Marketing",
	},

	"security-and-compliance": {
		"Cyber Security",
		"Legal/Tech Law",
	},
}
