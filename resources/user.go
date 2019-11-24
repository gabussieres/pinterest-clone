package resources

import (
	"fmt"

	"pinterest-clone/config"
	"pinterest-clone/models"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// UserInterface holds dynamo methods that can be performed on a user
type UserInterface interface {
	Get() (user *models.User, err error)
	QueryPins() (pinIDs Pins, err error)
}

// User is the string representation of a user ID. It implements the UserInterface
type User string

// Get returns the user for the given id
func (u User) Get() (*models.User, error) {
	return u.get(Dynamo{})
}

func (u User) get(dynamo DynamoInterface) (user *models.User, err error) {
	var result *dynamodb.GetItemOutput
	result, err = dynamo.GetItem(string(u), config.PinterestConfig.DynamoConfig.UserTable)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user = &models.User{}
	err = dynamodbattribute.UnmarshalMap(result.Item, user)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	return
}

// QueryPins returns an array of pinIDs for the user
func (u User) QueryPins() (Pins, error) {
	return u.queryPins(Dynamo{})
}

func (u User) queryPins(dynamo DynamoInterface) (pinIDs Pins, err error) {
	var result *dynamodb.QueryOutput
	result, err = dynamo.Query(string(u), config.PinterestConfig.DynamoConfig.UserPinsTable)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, item := range result.Items {
		id := item[config.PinterestConfig.DynamoConfig.UserPinsTable.SortKey].N
		pinIDs = append(pinIDs, Pin(*id))
	}
	return
}
