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
	testUserID     = "userID"
	testUserParams = pinterest.UserParams{
		UserID: testUserID,
	}
)

type fakeUserGetError struct{}

func (f fakeUserGetError) Get() (user *models.User, err error) {
	err = errors.New("Get error")
	return
}
func (f fakeUserGetError) QueryPins() (pinIDs resources.Pins, err error) {
	return
}

func TestHandleUserCallGetError(t *testing.T) {
	_, err := handleUserCall(fakeUserGetError{})
	if err == nil {
		t.Errorf("Expected an error. Got %v", err)
	}
}

type fakeUserGet struct{}

func (f fakeUserGet) Get() (user *models.User, err error) {
	user = &models.User{}
	return
}
func (f fakeUserGet) QueryPins() (pinIDs resources.Pins, err error) {
	return
}

func TestHandleUserCallHappyPath(t *testing.T) {
	user, err := handleUserCall(fakeUserGet{})
	if err != nil {
		t.Errorf("Expected no error. Got %v", err)
	}
	expectedUser := &models.User{}
	if !reflect.DeepEqual(expectedUser, user) {
		t.Errorf("Expected %v. Got %v", expectedUser, user)
	}
}
