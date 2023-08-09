package store

import "github.com/aws/aws-sdk-go/service/dynamodb"

var DB *dynamodb.DynamoDB

const (
	UsersTable = "checkmateafrica-users"
	PairsTable = "checkmateafrica-pairs"
)

// aws dynamodb create-table --endpoint-url http://localhost:8000 \
//     --table-name checkmateafrica-users \
//     --attribute-definitions \
//         AttributeName=Email,AttributeType=S \
//     --key-schema \
//         AttributeName=Email,KeyType=HASH \
//     --provisioned-throughput \
//         ReadCapacityUnits=5,WriteCapacityUnits=5 \
//     --table-class STANDARD

// dummyUser := User{
// 	Email:           "dummy@dummy.com",
// 	Name:            "John James",
// 	ExperienceLevel: "Beginner",
// 	SkillCategories: []string{"Frontend Development, Backend Development, Cloud Engineering, DevOps"},
// 	Gender:          "Male",
// 	SlackId:         "U01DYTEDHS8",
// }

/*
	Store options for partners
	- Save partner's user struct containing all details
		(id, name, skill category, etc) inside user struct in user's table
	- New table with single entry per rotation, rotation date as key and 'pairings' field
	 	pairings field is a map, user's email as keys, partner's user struct as value
	-	New table with single entry per rotation, rotation date as key and 'pairings' field
	 	pairings field is a slice of 'partners', and 'partners' is a slice of strings
		containing emails of both partners
	- Single entry per rotation, rotation date as key, map of all pairs,
		pair combined email as key, child-map of each user in the pair
		with user's email as key
	-	New table with one entry per pair, an entry consists of one partition key,
		and one global secondary index key, the keys are both partner's emails
		then last item is a map[email]userStruct
	- New table, pairId as partition key (0... n), two maps for each partner,
		user email as map key.
*/
