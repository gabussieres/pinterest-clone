// Code generated by go-swagger; DO NOT EDIT.

package pinterest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "pinterest-clone/models"
)

// PinOKCode is the HTTP code returned for type PinOK
const PinOKCode int = 200

/*PinOK a pin

swagger:response pinOK
*/
type PinOK struct {

	/*
	  In: Body
	*/
	Payload *models.Pin `json:"body,omitempty"`
}

// NewPinOK creates PinOK with default headers values
func NewPinOK() *PinOK {

	return &PinOK{}
}

// WithPayload adds the payload to the pin o k response
func (o *PinOK) WithPayload(payload *models.Pin) *PinOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pin o k response
func (o *PinOK) SetPayload(payload *models.Pin) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PinOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*PinDefault error

swagger:response pinDefault
*/
type PinDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPinDefault creates PinDefault with default headers values
func NewPinDefault(code int) *PinDefault {
	if code <= 0 {
		code = 500
	}

	return &PinDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the pin default response
func (o *PinDefault) WithStatusCode(code int) *PinDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the pin default response
func (o *PinDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the pin default response
func (o *PinDefault) WithPayload(payload *models.Error) *PinDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the pin default response
func (o *PinDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PinDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
