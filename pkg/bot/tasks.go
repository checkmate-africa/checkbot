package bot

import (
	"math"

	"github.com/checkmateafrica/accountability-bot/pkg/store"
	"github.com/checkmateafrica/accountability-bot/pkg/utils"
)

func GeneratePairs(users []store.User) [][]store.User {
	var pairs store.Pairs
	var pairedUsers = make(map[int]bool)

	/* create pairs of users with
	at least one matching skill category
	*/
	for i, userA := range users {
		if pairedUsers[i] {
			continue
		}

		for j := i + 1; j < len(users); j++ {
			if pairedUsers[j] {
				continue
			}

			userB := users[j]
			commonSkill := utils.FindCommonSkill(userA.SkillCategories, userB.SkillCategories)

			if commonSkill != "" {
				pairs = append(pairs, []store.User{userA, userB})
				pairedUsers[i] = true
				pairedUsers[j] = true
				break
			}
		}
	}

	/* if there are unpaired users left
	due to no matching skills, match
	using parent category
	*/
	if len(pairs) != int(math.Round(float64(len(users)/2)+0.5)) {
		for i, userA := range users {
			if pairedUsers[i] {
				continue
			}

			for j := i + 1; j < len(users); j++ {
				if pairedUsers[j] {
					continue
				}

				userB := users[j]
				commonParent := utils.FindCommonParent(userA.SkillCategories, userB.SkillCategories)

				if commonParent != "" {
					pairs = append(pairs, []store.User{userA, userB})
					pairedUsers[i] = true
					pairedUsers[j] = true
					break
				}
			}
		}
	}

	/* if there are still unpaired users left
	due to no matching parent category,
	pair randomly
	*/
	if len(pairs) != int(math.Round(float64(len(users)/2)+0.5)) {
		var unpairedUsers []store.User

		for i := range users {
			if !pairedUsers[i] {
				unpairedUsers = append(unpairedUsers, users[i])
			}
		}

		for i := 0; i < len(unpairedUsers); i += 2 {

			if i+1 < len(unpairedUsers) {
				pairs = append(pairs, []store.User{unpairedUsers[i], unpairedUsers[i+1]})
			} else {
				pairs = append(pairs, []store.User{unpairedUsers[i]})
			}
		}
	}

	return pairs
}
