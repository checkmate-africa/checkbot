package store

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetUsers() (*[]User, error) {
	users := []User{}

	result, err := DB.Scan(&dynamodb.ScanInput{
		TableName: aws.String(UsersTable),
	})

	if err != nil {
		log.Println(err)
		return nil, err
	}

	for _, item := range result.Items {
		user := User{}

		if err = dynamodbattribute.UnmarshalMap(item, &user); err != nil {
			log.Println(err)
		} else {
			users = append(users, user)
		}
	}

	return &users, nil
}

func GetUser(email string) (*User, error) {
	user := User{}

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
		return nil, err
	}

	if result.Item == nil {
		msg := "could not find user with email " + email
		log.Println(msg)

		return nil, err
	}

	if err = dynamodbattribute.UnmarshalMap(result.Item, &user); err != nil {
		log.Println(err)
	}

	return &user, nil
}

func SaveUserData(u User) error {
	user, err := dynamodbattribute.MarshalMap(u)

	if err != nil {
		log.Println(err)
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      user,
		TableName: aws.String(UsersTable),
	}

	_, err = DB.PutItem(input)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
