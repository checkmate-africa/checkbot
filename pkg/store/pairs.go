package store

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func SavePairs(pairs []Pair) {
	for _, pair := range pairs {
		go func(p Pair) {
			fmt.Println(p)
		}(pair)
	}
}

func GetPartner(email string) *User {
	result, err := DB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(PairsTable),
		Key: map[string]*dynamodb.AttributeValue{
			"Email": {
				S: aws.String(email),
			},
		},
	})

	if err != nil {
		log.Println(err)
	}

	if result.Item == nil {
		msg := "could not find partner for user " + email
		log.Println(msg)

		return nil
	}

	pair := Pair{}

	if err = dynamodbattribute.UnmarshalMap(result.Item, &pair); err != nil {
		log.Println(err)
	}

	return &pair.Partner
}
