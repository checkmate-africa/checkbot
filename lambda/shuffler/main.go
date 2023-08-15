package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/checkmateafrica/accountability-bot/pkg/bot"
	"github.com/checkmateafrica/accountability-bot/pkg/store"
	"github.com/checkmateafrica/accountability-bot/services"
)

func handler() error {
	users := *store.GetUsers()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range users {
		newPosition := r.Intn(len(users) - 1)
		users[i], users[newPosition] = users[newPosition], users[i]
	}

	pairs := bot.GeneratePairs(users)

	for _, pair := range pairs {
		marshaled, _ := json.MarshalIndent(pair, "", "   ")
		fmt.Println(string(marshaled))
		fmt.Println(" ")
	}

	return nil
}

func main() {
	log.SetPrefix("ERROR: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	store.DB = services.NewDynaClient()
	lambda.Start(handler)
}
