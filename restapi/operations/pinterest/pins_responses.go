// Code generated by go-swagger; DO NOT EDIT.

package pinterest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "pinterest-clone/models"
)

// PinsOKCode is the HTTP code returned for type PinsOK
const PinsOKCode int = 200

/*PinsOK the user's pins

swagger:response pinsOK
*/
type PinsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Pin `json:"body,omitempty"`
}

// NewPinsOK creates PinsOK with default headers values
func NewPinsOK() *PinsOK {

	return &PinsOK{}
}

// WithPayload adds the payload to the pins o k response
func (o *PinsOK) WithPayload(payload []*models.Pin) *PinsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pins o k response
func (o *PinsOK) SetPayload(payload []*models.Pin) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PinsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Pin, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*PinsDefault error

swagger:response pinsDefault
*/
type PinsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPinsDefault creates PinsDefault with default headers values
func NewPinsDefault(code int) *PinsDefault {
	if code <= 0 {
		code = 500
	}

	return &PinsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the pins default response
func (o *PinsDefault) WithStatusCode(code int) *PinsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the pins default response
func (o *PinsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the pins default response
func (o *PinsDefault) WithPayload(payload *models.Error) *PinsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pins default response
func (o *PinsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PinsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
