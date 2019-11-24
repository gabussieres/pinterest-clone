package controllers

import (
	"pinterest-clone/models"
	"pinterest-clone/resources"
	"pinterest-clone/restapi/operations/pinterest"
)

// HandlePinCall handles the /pin call, returns a pin for the given PinID
func HandlePinCall(params pinterest.PinParams) (pin *models.Pin, err error) {
	p := resources.Pin(params.PinID)
	return handlePinCall(p)
}

func handlePinCall(p resources.PinInterface) (pin *models.Pin, err error) {
	pin, err = p.Get()
	return
}
