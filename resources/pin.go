package resources

import (
	"fmt"

	"pinterest-clone/config"
	"pinterest-clone/models"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// PinsInterface holds methods that can be performed on an array of pins
type PinsInterface interface {
	Concatenate(pins Pins)
	BatchGet() (pinIDs []*models.Pin, err error)
}

// PinInterface holds methods that can be performed on a pin
type PinInterface interface {
	Get() (pin *models.Pin, err error)
}

// Pins implements the PinInterface
type Pins []Pin

// Pin is the string representation of a pin ID. It implements the PinInterface
type Pin string

// Get returns the pin for the given id
func (p Pin) Get() (*models.Pin, error) {
	return p.get(Dynamo{})
}

func (p Pin) get(dynamo DynamoInterface) (pin *models.Pin, err error) {
	var result *dynamodb.GetItemOutput
	result, err = dynamo.GetItem(string(p), config.PinterestConfig.DynamoConfig.PinTable)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	pin = &models.Pin{}
	err = dynamodbattribute.UnmarshalMap(result.Item, pin)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	return
}

// Concatenate loads new Pin IDs into Pins
func (p Pins) Concatenate(pins Pins) {
	p.concatenate(pins)
}

func (p *Pins) concatenate(pins Pins) {
	(*p) = append(*p, pins...)
}

// BatchGet returns an array of pins for the given pin ids
func (p Pins) BatchGet() ([]*models.Pin, error) {
	return p.batchGet(Dynamo{})
}

func (p Pins) batchGet(dynamo DynamoInterface) (pins []*models.Pin, err error) {
	var result *dynamodb.BatchGetItemOutput
	result, err = dynamo.BatchGetItem(p, config.PinterestConfig.DynamoConfig.PinTable)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, p := range result.Responses[config.PinterestConfig.DynamoConfig.PinTable.Name] {
		pin := &models.Pin{}
		err = dynamodbattribute.UnmarshalMap(p, pin)
		if err != nil {
			panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
		}
		pins = append(pins, pin)
	}
	return
}
