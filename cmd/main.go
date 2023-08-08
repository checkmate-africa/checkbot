package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/checkmateafrica/users/pkg/handlers"
	"github.com/checkmateafrica/users/pkg/store"
	"github.com/gorilla/mux"
)

func main() {
	log.SetPrefix("ERROR: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	awsSess, err := session.NewSession(&aws.Config{
		Region:      aws.String("local"),
		Credentials: credentials.NewStaticCredentials("x", "x", ""),
		Endpoint:    aws.String("http://localhost:8000"),
	})

	if err != nil {
		log.Println(err)
		return
	}

	dynaClient := dynamodb.New(awsSess)
	store.DB = dynaClient

	port := ":4000"
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handlers.DefaultHandler).Methods("GET")
	router.HandleFunc("/event", handlers.EventsHandler).Methods("POST")
	router.HandleFunc("/interaction", handlers.InteractionsHandler).Methods("POST")

	fmt.Println("listening on port", port)

	log.Fatal(http.ListenAndServe(port, router))

	// lambda.Start(httpadapter.New(router).ProxyWithContext)
}
