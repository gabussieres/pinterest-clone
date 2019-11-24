package controllers

import (
	"pinterest-clone/models"
	"pinterest-clone/resources"
	"pinterest-clone/restapi/operations/pinterest"
)

// HandleUserCall handles the /user call, returns the user for the given UserID
func HandleUserCall(params pinterest.UserParams) (user *models.User, err error) {
	u := resources.User(params.UserID)
	return handleUserCall(u)
}

func handleUserCall(u resources.UserInterface) (user *models.User, err error) {
	user, err = u.Get()
	return
}
