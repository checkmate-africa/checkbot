package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/checkmateafrica/accountability-bot/internal/store"
)

func IsLocalEnv() bool {
	return os.Getenv(EnvSamLocal) == "true"
}

func GetNextDayOfWeek(dayOfWeek int) (day string, month string, year string, daysUntil string) {
	var (
		CurrentDate   = time.Now()
		DaysUntilNext = (dayOfWeek + 7 - int(CurrentDate.Weekday())) % 7
		NextDate      = CurrentDate.AddDate(0, 0, DaysUntilNext+7) // extra 7 days
	)

	y, m, d := NextDate.Date()

	return strconv.Itoa(d), m.String(), strconv.Itoa(y), strconv.Itoa(DaysUntilNext + 7)
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

func FindCurrentPartner(pairs []store.PairedUser, email string) *store.User {
	for _, pair := range pairs {
		if pair.Email == email {
			return &pair.Partner
		}
	}
	return nil
}
