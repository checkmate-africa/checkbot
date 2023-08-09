package bot

import "github.com/checkmateafrica/users/pkg/store"

func GenerateAccountabilityPairs() {
	var pairs []store.Pair

	store.SavePairs(pairs)
}

func CreateFocusRooms() {
	// google calender api?
	// dyte community access?
}
