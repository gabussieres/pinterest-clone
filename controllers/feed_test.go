package controllers

import (
	"errors"
	"pinterest-clone/models"
	"pinterest-clone/resources"
	"reflect"
	"testing"
)

type fakeFeedBatchGetError struct{}

func (f fakeFeedBatchGetError) BatchGet() (pinIDs []*models.Pin, err error) {
	err = errors.New("BatchGet error")
	return
}
func (f fakeFeedBatchGetError) Concatenate(pins resources.Pins) {
	return
}

func TestHandleFeedCallBatchGetError(t *testing.T) {
	_, err := handleFeedCall(fakeFeedBatchGetError{})
	if err == nil {
		t.Errorf("Expected an error. Got %v", err)
	}
}

type fakeFeed struct{}

func (f fakeFeed) Concatenate(pins resources.Pins) {
	pins = resources.Pins{testPin1, testPin2}
	return
}
func (f fakeFeed) BatchGet() (pinIDs []*models.Pin, err error) {
	pinIDs = testPins
	return
}

func TestHandleFeedCallHappyPath(t *testing.T) {
	pins, err := handleFeedCall(fakeFeed{})
	if err != nil {
		t.Errorf("Expected no error. Got %v", err)
	}
	if !reflect.DeepEqual(testPins, pins) {
		t.Errorf("Expected %v. Got %v", testPins, pins)
	}
}
