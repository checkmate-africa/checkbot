package store

import "github.com/aws/aws-sdk-go/service/dynamodb"

var DB *dynamodb.DynamoDB
var TableName string = "checkmateafrica-users"

// aws dynamodb create-table --endpoint-url http://localhost:8000 \
//     --table-name checkmateafrica-users \
//     --attribute-definitions \
//         AttributeName=Email,AttributeType=S \
//     --key-schema \
//         AttributeName=Email,KeyType=HASH \
//     --provisioned-throughput \
//         ReadCapacityUnits=5,WriteCapacityUnits=5 \
//     --table-class STANDARD
