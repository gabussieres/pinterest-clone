package controllers

import (
	"math/rand"
	"pinterest-clone/models"
	"pinterest-clone/resources"
	"strconv"
)

// HandleFeedCall handles the /feed request, returns a random array of pins
func HandleFeedCall() (feed []*models.Pin, err error) {
	pins := resources.Pins{}
	return handleFeedCall(pins)
}

func handleFeedCall(p resources.PinsInterface) (feed []*models.Pin, err error) {
	var pins resources.Pins
	for i := 1; i < 301; i++ {
		pins = append(pins, resources.Pin(strconv.Itoa(i)))
	}
	rand.Shuffle(len(pins), func(i, j int) {
		pins[i], pins[j] = pins[j], pins[i]
	})
	p.Concatenate(pins)
	feed, err = p.BatchGet()
	return
}
