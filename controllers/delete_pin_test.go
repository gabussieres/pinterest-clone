package controllers

import (
	"errors"
	"pinterest-clone/models"
	"pinterest-clone/resources"
	"pinterest-clone/restapi/operations/pinterest"
	"testing"
)

var (
	testDeletePinParams = pinterest.DeletePinParams{
		UserID: testUserID,
		PinID:  testPinID,
	}
)

type fakeUserDeletePinError struct{}

func (f fakeUserDeletePinError) Get() (user *models.User, err error) {
	return
}
func (f fakeUserDeletePinError) QueryPins() (pinIDs resources.Pins, err error) {
	return
}
func (f fakeUserDeletePinError) DeletePin(pinID string) (err error) {
	err = errors.New("DeletePin error")
	return
}

func TestHandleUserCallDeletePinError(t *testing.T) {
	err := handleDeletePinCall(fakeUserDeletePinError{}, testPinID)
	if err == nil {
		t.Errorf("Expected an error. Got %v", err)
	}
}

type fakeUserDeletePin struct{}

func (f fakeUserDeletePin) Get() (user *models.User, err error) {
	return
}
func (f fakeUserDeletePin) QueryPins() (pinIDs resources.Pins, err error) {
	return
}
func (f fakeUserDeletePin) DeletePin(pinID string) (err error) {
	return
}

func TestHandleDeletePinCallHappyPath(t *testing.T) {
	err := handleDeletePinCall(fakeUserDeletePin{}, testPinID)
	if err != nil {
		t.Errorf("Expected no error. Got %v", err)
	}
}
