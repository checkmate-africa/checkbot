package store

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func GetUsers() {

}

func SaveUserData(u User) {
	user, err := dynamodbattribute.MarshalMap(u)

	if err != nil {
		fmt.Println("Marshal Error: ", err)
		return
	}

	input := &dynamodb.PutItemInput{
		Item:      user,
		TableName: aws.String(TableName),
	}

	_, err = DB.PutItem(input)

	if err != nil {
		fmt.Println("Marshal Error: ", err)
		return
	}

	fmt.Println(user)
}
