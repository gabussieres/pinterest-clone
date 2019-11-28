package controllers

import (
	"pinterest-clone/models"
	"pinterest-clone/resources"
	"pinterest-clone/restapi/operations/pinterest"
)

// HandlePinsCall handles the /pins call, returns an array of pins for the given UserID
func HandlePinsCall(params pinterest.PinsParams) ([]*models.Pin, error) {
	user := resources.User(params.UserID)
	pins := &resources.Pins{}
	return handlePinsCall(params, user, pins)
}

func handlePinsCall(params pinterest.PinsParams, u resources.UserInterface, p resources.PinsInterface) (pins []*models.Pin, err error) {
	var queriedPins resources.Pins
	queriedPins, err = u.QueryPins()
	if err != nil {
		return
	}
	p.Concatenate(queriedPins)
	pins, err = p.BatchGet()
	return
}
