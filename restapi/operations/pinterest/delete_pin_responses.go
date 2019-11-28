// Code generated by go-swagger; DO NOT EDIT.

package pinterest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "pinterest-clone/models"
)

// DeletePinOKCode is the HTTP code returned for type DeletePinOK
const DeletePinOKCode int = 200

/*DeletePinOK pin is deleted

swagger:response deletePinOK
*/
type DeletePinOK struct {
}

// NewDeletePinOK creates DeletePinOK with default headers values
func NewDeletePinOK() *DeletePinOK {

	return &DeletePinOK{}
}

// WriteResponse to the client
func (o *DeletePinOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*DeletePinDefault error

swagger:response deletePinDefault
*/
type DeletePinDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeletePinDefault creates DeletePinDefault with default headers values
func NewDeletePinDefault(code int) *DeletePinDefault {
	if code <= 0 {
		code = 500
	}

	return &DeletePinDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete pin default response
func (o *DeletePinDefault) WithStatusCode(code int) *DeletePinDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete pin default response
func (o *DeletePinDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete pin default response
func (o *DeletePinDefault) WithPayload(payload *models.Error) *DeletePinDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete pin default response
func (o *DeletePinDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePinDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}