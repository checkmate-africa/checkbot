package bot

import (
	"github.com/checkmateafrica/accountability-bot/pkg/store"
	"github.com/checkmateafrica/accountability-bot/pkg/utils"
)

func GenerateAccountabilityPairs() {
	var pairs []store.Pair

	// fetch users
	// run algorithm
	// send message

	store.SavePairs(pairs)
}

func GeneratePairs(users []store.User) [][]store.User {
	var pairs [][]store.User

	for i, userA := range users {
		for j := i + 1; j < len(users); j++ {
			userB := users[j]
			commonSkill := findCommonSkill(userA.SkillCategories, userB.SkillCategories)

			if commonSkill != "" {
				// update both users in database
				pairs = append(pairs, []store.User{userA, userB})
			}
		}
	}

	// If no common skill pairs found, try same parent-category
	if len(pairs) == 0 {
		for i, userA := range users {
			for j := i + 1; j < len(users); j++ {
				userB := users[j]
				commonParent := findCommonParent(userA.SkillCategories, userB.SkillCategories)

				if commonParent != "" {
					// update both users in database
					pairs = append(pairs, []store.User{userA, userB})
				}
			}
		}
	}

	// If still no pairs found, match users randomly
	if len(pairs) == 0 {
		for i := 0; i < len(users); i += 2 {
			if i+1 < len(users) {
				// update both users in database
				pairs = append(pairs, []store.User{users[i], users[i+1]})
			} else {
				// Handle odd number of users
				// update single user in database
				pairs = append(pairs, []store.User{users[i]})
			}
		}
	}

	return pairs
}

func findCommonSkill(skillsA, skillsB []string) string {
	for _, skillA := range skillsA {
		for _, skillB := range skillsB {
			if skillA == skillB {
				return skillA
			}
		}
	}
	return ""
}

func findCommonParent(skillsA, skillsB []string) string {
	for _, skillA := range skillsA {
		for _, skillB := range skillsB {
			parentA := findParentCategory(skillA)
			parentB := findParentCategory(skillB)
			if parentA != "" && parentA == parentB {
				return parentA
			}
		}
	}
	return ""
}

func findParentCategory(skill string) string {
	for parent, skills := range utils.SkillDomains {
		for _, s := range skills {
			if s == skill {
				return parent
			}
		}
	}
	return ""
}

func CreateFocusRooms() {
	// google calender api?
	// dyte community access?
}
