package controllers

import (
	"math/rand"
	"strconv"

	"pinterest-clone/models"
	"pinterest-clone/resources"
	"pinterest-clone/restapi/operations/pinterest"
)

// HandleFeedCall handles the /feed request, returns a random array of pins
func HandleFeedCall(params pinterest.FeedParams) (feed []*models.Pin, err error) {
	pins := &resources.Pins{}
	return handleFeedCall(pins, params)
}

func handleFeedCall(p resources.PinsInterface, params pinterest.FeedParams) (feed []*models.Pin, err error) {
	var pins resources.Pins
	for i := 1; i < 301; i++ {
		pins = append(pins, resources.Pin(strconv.Itoa(i)))
	}
	rand.Shuffle(len(pins), func(i, j int) {
		pins[i], pins[j] = pins[j], pins[i]
	})

	var pinAmount int
	pinAmount, err = strconv.Atoi(params.PinAmount)
	if err != nil {
		return
	}
	p.Concatenate(pins[:pinAmount])
	feed, err = p.BatchGet()
	return
}
