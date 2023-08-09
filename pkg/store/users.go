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
