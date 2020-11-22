// Code generated by go-swagger; DO NOT EDIT.

package stock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gamersnote.com/v1/models"
)

// DeleteStockNoContentCode is the HTTP code returned for type DeleteStockNoContent
const DeleteStockNoContentCode int = 204

/*DeleteStockNoContent Deleted

swagger:response deleteStockNoContent
*/
type DeleteStockNoContent struct {
}

// NewDeleteStockNoContent creates DeleteStockNoContent with default headers values
func NewDeleteStockNoContent() *DeleteStockNoContent {

	return &DeleteStockNoContent{}
}

// WriteResponse to the client
func (o *DeleteStockNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*DeleteStockDefault error

swagger:response deleteStockDefault
*/
type DeleteStockDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteStockDefault creates DeleteStockDefault with default headers values
func NewDeleteStockDefault(code int) *DeleteStockDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteStockDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete stock default response
func (o *DeleteStockDefault) WithStatusCode(code int) *DeleteStockDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete stock default response
func (o *DeleteStockDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete stock default response
func (o *DeleteStockDefault) WithPayload(payload *models.Error) *DeleteStockDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete stock default response
func (o *DeleteStockDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteStockDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
