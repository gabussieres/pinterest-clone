package controllers

import (
	"pinterest-clone/resources"
	"pinterest-clone/restapi/operations/pinterest"
)

// HandleDeletePinCall handles the DELETE /pin call, deletes the pin for the given PinID
func HandleDeletePinCall(params pinterest.DeletePinParams) error {
	u := resources.User(params.UserID)
	return handleDeletePinCall(u, params.PinID)
}

func handleDeletePinCall(u resources.UserInterface, pinID string) (err error) {
	err = u.DeletePin(pinID)
	return
}
