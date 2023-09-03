package store

import (
	"log"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func SavePairs(pairs Pairs) []PairedUser {
	pairedUsers := []PairedUser{}
	var wg sync.WaitGroup

	for _, p := range pairs {
		wg.Add(1)

		go func(pair []User) {
			defer wg.Done()

			if len(pair) > 1 {
				for i, user := range pair {
					parterIndex := map[int]int{
						0: 1,
						1: 0,
					}

					pairedUser := PairedUser{
						Email:   user.Email,
						SlackId: user.SlackId,
						Partner: pair[parterIndex[i]],
					}

					pairedUserMap, err := dynamodbattribute.MarshalMap(pairedUser)

					if err != nil {
						log.Println(err)
						continue
					}

					input := &dynamodb.PutItemInput{
						Item:      pairedUserMap,
						TableName: aws.String(PairsTable),
					}

					_, err = DB.PutItem(input)

					if err != nil {
						log.Println(err)
						continue
					}

					pairedUsers = append(pairedUsers, pairedUser)
				}
			} else {
				_, err := DB.DeleteItem(&dynamodb.DeleteItemInput{
					TableName: aws.String(PairsTable),
					Key: map[string]*dynamodb.AttributeValue{
						"Email": {
							S: aws.String(pair[0].Email),
						},
					},
				})

				if err != nil {
					log.Println(err)
				}
			}
		}(p)
	}

	wg.Wait()

	return pairedUsers
}

func GetPartner(email string) (*User, error) {
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
		log.Println("could not find partner for user " + email)

		return nil, err
	}

	var pair PairedUser

	if err = dynamodbattribute.UnmarshalMap(result.Item, &pair); err != nil {
		log.Println(err)
	}

	return &pair.Partner, nil
}
