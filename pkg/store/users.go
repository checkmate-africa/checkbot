package store

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetUsers() {

}

func GetUser(email string) *User {
	result, err := DB.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(UsersTable),
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
		msg := "could not user with email " + email
		log.Println(msg)

		return nil
	}

	user := User{}

	if err = dynamodbattribute.UnmarshalMap(result.Item, &user); err != nil {
		log.Println(err)
	}

	return &user
}

func SaveUserData(u User) {
	user, err := dynamodbattribute.MarshalMap(u)

	if err != nil {
		log.Println(err)
		return
	}

	input := &dynamodb.PutItemInput{
		Item:      user,
		TableName: aws.String(UsersTable),
	}

	_, err = DB.PutItem(input)

	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(user)
}
