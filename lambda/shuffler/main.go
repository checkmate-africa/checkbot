package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/checkmateafrica/accountability-bot/internal/bot"
	"github.com/checkmateafrica/accountability-bot/internal/store"
	"github.com/checkmateafrica/accountability-bot/services"
)

func handler() error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	u, err := store.GetUsers()

	if err != nil {
		log.Println(err)
		return err
	}

	users := *u

	for i := range users {
		newPosition := r.Intn(len(users) - 1)
		users[i], users[newPosition] = users[newPosition], users[i]
	}

	pairs := bot.GeneratePairs(users)
	pairedUsers := store.SavePairs(pairs)

	if err = bot.SendNewPairShuffleAnnouncement(pairs); err != nil {
		log.Println(err)
		return err
	}

	bot.SendPairShuffleNotification(pairedUsers)

	return nil
}

func main() {
	log.SetPrefix("ERROR: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	store.DB = services.NewDynaClient()
	lambda.Start(handler)
}
