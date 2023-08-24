package store

import "github.com/aws/aws-sdk-go/service/dynamodb"

var DB *dynamodb.DynamoDB

const (
	UsersTable = "checkmateafrica-users"
	PairsTable = "checkmateafrica-pairs"
)
