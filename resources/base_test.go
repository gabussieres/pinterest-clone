package resources

import (
	"errors"
	"reflect"
	"testing"

	"pinterest-clone/config"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	testID1         = "id1"
	testID2         = "id2"
	testTableConfig = config.TableConfig{
		Name:         "table",
		PartitionKey: "pk",
		SortKey:      "sk",
	}
)

func createString(s string) *string {
	return &s
}

type fakeDBClientError struct{ dynamodbiface.DynamoDBAPI }

func (f fakeDBClientError) GetItem(in *dynamodb.GetItemInput) (out *dynamodb.GetItemOutput, err error) {
	err = errors.New("GetItemError")
	return
}

func TestGetError(t *testing.T) {
	d := Dynamo{}
	_, err := d.getItem(&fakeDBClientError{}, testID1, testTableConfig)
	if err == nil {
		t.Errorf("Expected an error. Got %v", err)
	}
}

type fakeDBClientEmptyResult struct{ dynamodbiface.DynamoDBAPI }

func (f fakeDBClientEmptyResult) GetItem(in *dynamodb.GetItemInput) (out *dynamodb.GetItemOutput, err error) {
	out = &dynamodb.GetItemOutput{}
	return
}
func (f fakeDBClientEmptyResult) BatchGetItem(in *dynamodb.BatchGetItemInput) (out *dynamodb.BatchGetItemOutput, err error) {
	out = &dynamodb.BatchGetItemOutput{}
	return
}
func (f fakeDBClientEmptyResult) Query(in *dynamodb.QueryInput) (out *dynamodb.QueryOutput, err error) {
	out = &dynamodb.QueryOutput{}
	return
}

func TestGetEmptyResult(t *testing.T) {
	d := Dynamo{}
	res, err := d.getItem(&fakeDBClientEmptyResult{}, testID1, testTableConfig)
	if err != nil {
		t.Errorf("Expected no error. Got %v", err)
	}
	expectedResult := &dynamodb.GetItemOutput{}
	if !reflect.DeepEqual(res, expectedResult) {
		t.Errorf("Expected %v. Got %v", expectedResult, res)
	}
}

func TestBatchGetEmptyResult(t *testing.T) {
	d := Dynamo{}
	res, err := d.batchGetItem(&fakeDBClientEmptyResult{}, testPins, testTableConfig)
	if err != nil {
		t.Errorf("Expected no error. Got %v", err)
	}
	expectedResult := &dynamodb.BatchGetItemOutput{}
	if !reflect.DeepEqual(res, expectedResult) {
		t.Errorf("Expected %v. Got %v", expectedResult, res)
	}
}

func TestQueryEmptyResult(t *testing.T) {
	d := Dynamo{}
	res, err := d.query(&fakeDBClientEmptyResult{}, testID1, testTableConfig)
	if err != nil {
		t.Errorf("Expected no error. Got %v", err)
	}
	expectedResult := &dynamodb.QueryOutput{}
	if !reflect.DeepEqual(res, expectedResult) {
		t.Errorf("Expected %v. Got %v", expectedResult, res)
	}
}
