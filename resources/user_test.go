package resources

import (
	"errors"
	"pinterest-clone/config"
	"pinterest-clone/models"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	testUserID = "userID"
	testUser   = User(testUserID)
)

type fakeDBClientUserError struct{ DynamoInterface }

func (f fakeDBClientUserError) GetItem(id string, tableConfig config.TableConfig) (out *dynamodb.GetItemOutput, err error) {
	err = errors.New("Dynamo.GetItem error")
	return
}

func TestGetUserError(t *testing.T) {
	_, err := testUser.get(&fakeDBClientUserError{})
	if err == nil {
		t.Errorf("Expected no error. Got %v", err)
	}
}

type fakeDBClientEmptyUser struct{ DynamoInterface }

func (f fakeDBClientEmptyUser) GetItem(id string, tableConfig config.TableConfig) (out *dynamodb.GetItemOutput, err error) {
	out = &dynamodb.GetItemOutput{}
	return
}

func TestGetUserEmptyUser(t *testing.T) {
	user, err := testUser.get(&fakeDBClientEmptyUser{})
	if err != nil {
		t.Errorf("Expected no error. Got %v", err)
	}
	expectedUser := &models.User{}
	if !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("Expected %v. Got %v", expectedUser, user)
	}
}

type fakeDBClientUserHappyPath struct{ DynamoInterface }

func (f fakeDBClientUserHappyPath) GetItem(id string, tableConfig config.TableConfig) (out *dynamodb.GetItemOutput, err error) {
	out = &dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"name": {
				S: &testUserID,
			},
			"image_url": {
				S: createString("image_url"),
			},
			"following": {
				N: createString("1"),
			},
			"followers": {
				N: createString("2"),
			},
		},
	}
	return
}

func TestGetUserHappyPath(t *testing.T) {
	user, err := testUser.get(&fakeDBClientUserHappyPath{})
	if err != nil {
		t.Errorf("Expected no error. Got %v", err)
	}
	expectedUser := &models.User{Name: testUserID, ImageURL: "image_url", Following: 1, Followers: 2}
	if !reflect.DeepEqual(user, expectedUser) {
		t.Errorf("Expected %v. Got %v", expectedUser, user)
	}
}

type fakeDBClientUserPinsError struct{ DynamoInterface }

func (f fakeDBClientUserPinsError) Query(id string, tableConfig config.TableConfig) (out *dynamodb.QueryOutput, err error) {
	err = errors.New("Dynamo.Query error")
	return
}

func TestQueryUserPinsError(t *testing.T) {
	_, err := testUser.queryPins(&fakeDBClientUserPinsError{})
	if err == nil {
		t.Errorf("Expected no error. Got %v", err)
	}
}

type fakeDBClientUserPinsHappyPath struct{ DynamoInterface }

func (f fakeDBClientUserPinsHappyPath) Query(id string, tableConfig config.TableConfig) (out *dynamodb.QueryOutput, err error) {
	out = &dynamodb.QueryOutput{
		Items: []map[string]*dynamodb.AttributeValue{{
			"pin_id": {
				S: &testPinID1,
			},
		}, {
			"pin_id": {
				S: &testPinID2,
			},
		}},
	}
	return
}

func TestQueryUserPinsHappyPath(t *testing.T) {
	pins, err := testUser.queryPins(&fakeDBClientUserPinsHappyPath{})
	if err != nil {
		t.Errorf("Expected no error. Got %v", err)
	}
	expectedPinIDs := testPins
	if !reflect.DeepEqual(expectedPinIDs, pins) {
		t.Errorf("Expected %v. Got %v", expectedPinIDs, pins)
	}
}
