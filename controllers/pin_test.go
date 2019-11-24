package controllers

import (
	"errors"
	"pinterest-clone/models"
	"pinterest-clone/restapi/operations/pinterest"
	"reflect"
	"testing"
)

var (
	testPinID     = "pinID"
	testPinParams = pinterest.PinParams{
		PinID: testPinID,
	}
)

type fakePinGetError struct{}

func (f fakePinGetError) Get() (pin *models.Pin, err error) {
	err = errors.New("Get error")
	return
}

func TestHandlePinCallGetError(t *testing.T) {
	_, err := handlePinCall(fakePinGetError{})
	if err == nil {
		t.Errorf("Expected an error. Got %v", err)
	}
}

type fakePinGet struct{}

func (f fakePinGet) Get() (pin *models.Pin, err error) {
	pin = &models.Pin{}
	return
}

func TestHandlePinCallHappyPath(t *testing.T) {
	pin, err := handlePinCall(fakePinGet{})
	if err != nil {
		t.Errorf("Expected no error. Got %v", err)
	}
	expectedPin := &models.Pin{}
	if !reflect.DeepEqual(expectedPin, pin) {
		t.Errorf("Expected %v. Got %v", expectedPin, pin)
	}
}
