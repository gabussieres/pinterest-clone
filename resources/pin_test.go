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
	testPinID1 = "1"
	testPinID2 = "2"
	testPin1   = Pin("1")
	testPin2   = Pin("2")
	testPins   = Pins{testPin1, testPin2}
)

type fakeDBClientPinError struct{ DynamoInterface }

func (f fakeDBClientPinError) GetItem(id string, tableConfig config.TableConfig) (out *dynamodb.GetItemOutput, err error) {
	err = errors.New("Dynamo.GetItem error")
	return
}
func (f fakeDBClientPinError) BatchGetItem(id interface{}, tableConfig config.TableConfig) (out *dynamodb.BatchGetItemOutput, err error) {
	err = errors.New("Dynamo.BatchGetItem error")
	return
}

func TestGetPinError(t *testing.T) {
	_, err := testPin1.get(&fakeDBClientPinError{})
	if err == nil {
		t.Errorf("Expected no error. Got %v", err)
	}
}
func TestBatchGetPinsError(t *testing.T) {
	_, err := testPins.batchGet(&fakeDBClientPinError{})
	if err == nil {
		t.Errorf("Expected no error. Got %v", err)
	}
}

type fakeDBClientEmptyPin struct{ DynamoInterface }

func (f fakeDBClientEmptyPin) GetItem(id string, tableConfig config.TableConfig) (out *dynamodb.GetItemOutput, err error) {
	out = &dynamodb.GetItemOutput{}
	return
}

func TestGetPinEmptyPin(t *testing.T) {
	pin, err := testPin1.get(&fakeDBClientEmptyPin{})
	if err != nil {
		t.Errorf("Expected no error. Got %v", err)
	}
	expectedPin := &models.Pin{}
	if !reflect.DeepEqual(pin, expectedPin) {
		t.Errorf("Expected %v. Got %v", expectedPin, pin)
	}
}

type fakeDBClientPinHappyPath struct{ DynamoInterface }

func (f fakeDBClientPinHappyPath) GetItem(id string, tableConfig config.TableConfig) (out *dynamodb.GetItemOutput, err error) {
	out = &dynamodb.GetItemOutput{
		Item: map[string]*dynamodb.AttributeValue{
			"title": {
				S: &testPinID1,
			},
			"description": {
				S: createString("description"),
			},
			"image_url": {
				S: createString("image_url"),
			},
			"source_url": {
				S: createString("source_url"),
			},
			"posted_by": {
				S: createString("posted_by"),
			},
		},
	}
	return
}
func (f fakeDBClientPinHappyPath) BatchGetItem(id interface{}, tableConfig config.TableConfig) (out *dynamodb.BatchGetItemOutput, err error) {
	out = &dynamodb.BatchGetItemOutput{
		Responses: map[string][]map[string]*dynamodb.AttributeValue{
			"pinterest_pins": []map[string]*dynamodb.AttributeValue{{
				"title": {
					S: &testPinID1,
				},
				"description": {
					S: createString("description"),
				},
				"image_url": {
					S: createString("image_url"),
				},
				"source_url": {
					S: createString("source_url"),
				},
				"posted_by": {
					S: createString("posted_by"),
				},
			}, {
				"title": {
					S: &testPinID2,
				},
				"description": {
					S: createString("description"),
				},
				"image_url": {
					S: createString("image_url"),
				},
				"source_url": {
					S: createString("source_url"),
				},
				"posted_by": {
					S: createString("posted_by"),
				},
			}},
		},
	}
	return
}

func TestGetPinHappyPath(t *testing.T) {
	pin, err := testPin1.get(&fakeDBClientPinHappyPath{})
	if err != nil {
		t.Errorf("Expected no error. Got %v", err)
	}
	expectedPin := &models.Pin{Title: testPinID1, Description: "description", ImageURL: "image_url", SourceURL: "source_url", PostedBy: "posted_by"}
	if !reflect.DeepEqual(pin, expectedPin) {
		t.Errorf("Expected %v. Got %v", expectedPin, pin)
	}
}

func TestBatchGetPinHappyPath(t *testing.T) {
	pins, err := testPins.batchGet(&fakeDBClientPinHappyPath{})
	if err != nil {
		t.Errorf("Expected no error. Got %v", err)
	}
	expectedPin1 := &models.Pin{Title: testPinID1, Description: "description", ImageURL: "image_url", SourceURL: "source_url", PostedBy: "posted_by"}
	expectedPin2 := &models.Pin{Title: testPinID2, Description: "description", ImageURL: "image_url", SourceURL: "source_url", PostedBy: "posted_by"}
	if !reflect.DeepEqual(pins[0], expectedPin1) {
		t.Errorf("Expected %v. Got %v", expectedPin1, pins[0])
	}
	if !reflect.DeepEqual(pins[1], expectedPin2) {
		t.Errorf("Expected %v. Got %v", expectedPin2, pins[1])
	}
}

func TestConcatenate(t *testing.T) {
	pins := Pins{testPin1}
	pinsToConcatenate := testPin2
	pins.concatenate(Pins{pinsToConcatenate})
	if !reflect.DeepEqual(testPins, pins) {
		t.Errorf("Expected %v. Got %v", testPins, pins)
	}
}
