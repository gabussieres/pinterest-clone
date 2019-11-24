package controllers

import (
	"errors"
	"pinterest-clone/models"
	"pinterest-clone/resources"
	"pinterest-clone/restapi/operations/pinterest"
	"reflect"
	"testing"
)

var (
	testPinsParams = pinterest.PinsParams{
		UserID: testUserID,
	}
	testPin1      = resources.Pin("1")
	testPin2      = resources.Pin("2")
	testPinModel1 = models.Pin{}
	testPinModel2 = models.Pin{}
	testPins      = []*models.Pin{&testPinModel1, &testPinModel2}
)

type fakePinsQueryError struct{}

func (f fakePinsQueryError) QueryPins() (pinIDs resources.Pins, err error) {
	err = errors.New("Get error")
	return
}
func (f fakePinsQueryError) Get() (user *models.User, err error) {
	return
}

type fakePinsBatchGetNull struct{}

func (f fakePinsBatchGetNull) BatchGet() (pinIDs []*models.Pin, err error) {
	return
}
func (f fakePinsBatchGetNull) Concatenate(pins resources.Pins) {
	return
}

func TestHandlePinsCallQueryError(t *testing.T) {
	_, err := handlePinsCall(testPinsParams, fakePinsQueryError{}, fakePinsBatchGetNull{})
	if err == nil {
		t.Errorf("Expected an error. Got %v", err)
	}
}

type fakePinsQueryGetNull struct{}

func (f fakePinsQueryGetNull) QueryPins() (pinIDs resources.Pins, err error) {
	return
}
func (f fakePinsQueryGetNull) Get() (user *models.User, err error) {
	return
}

type fakePinsBatchGetError struct{}

func (f fakePinsBatchGetError) BatchGet() (pinIDs []*models.Pin, err error) {
	err = errors.New("BatchGet error")
	return
}
func (f fakePinsBatchGetError) Concatenate(pins resources.Pins) {
	return
}

func TestHandlePinsCallBatchGetError(t *testing.T) {
	_, err := handlePinsCall(testPinsParams, fakePinsQueryGetNull{}, fakePinsBatchGetError{})
	if err == nil {
		t.Errorf("Expected an error. Got %v", err)
	}
}

type fakePinsQuery struct{}

func (f fakePinsQuery) QueryPins() (pinIDs resources.Pins, err error) {
	pinIDs = resources.Pins{testPin1, testPin2}
	return
}
func (f fakePinsQuery) Get() (user *models.User, err error) {
	return
}

type fakePinsBatchGet struct{}

func (f fakePinsBatchGet) Concatenate(pins resources.Pins) {
	pins = resources.Pins{testPin1, testPin2}
	return
}
func (f fakePinsBatchGet) BatchGet() (pinIDs []*models.Pin, err error) {
	pinIDs = testPins
	return
}

func TestHandlePinsCallHappyPath(t *testing.T) {
	pins, err := handlePinsCall(testPinsParams, fakePinsQuery{}, fakePinsBatchGet{})
	if err != nil {
		t.Errorf("Expected no error. Got %v", err)
	}
	if !reflect.DeepEqual(testPins, pins) {
		t.Errorf("Expected %v. Got %v", testPins, pins)
	}
}
