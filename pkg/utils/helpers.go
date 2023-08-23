package utils

import (
	"os"
	"strconv"
	"time"
)

func IsLocalEnv() bool {
	return os.Getenv(EnvSamLocal) == "true"
}

func GetNextDayOfWeek(dayOfWeek int) (day string, month string, year string, daysUntil string) {
	var (
		CurrentDate   = time.Now()
		DaysUntilNext = (dayOfWeek + 7 - int(CurrentDate.Weekday())) % 7
		NextDate      = CurrentDate.AddDate(0, 0, DaysUntilNext)
	)

	y, m, d := NextDate.Date()

	return strconv.Itoa(d), m.String(), strconv.Itoa(y), strconv.Itoa(DaysUntilNext)
}

func FindCommonSkill(skillsA, skillsB []string) string {
	for _, skillA := range skillsA {
		for _, skillB := range skillsB {
			if skillA == skillB {
				return skillA
			}
		}
	}
	return ""
}

func FindCommonParent(skillsA, skillsB []string) string {
	for _, skillA := range skillsA {
		for _, skillB := range skillsB {
			parentA := FindParentCategory(skillA)
			parentB := FindParentCategory(skillB)
			if parentA != "" && parentA == parentB {
				return parentA
			}
		}
	}
	return ""
}

func FindParentCategory(skill string) string {
	for parent, skills := range SkillDomains {
		for _, s := range skills {
			if s == skill {
				return parent
			}
		}
	}
	return ""
}
