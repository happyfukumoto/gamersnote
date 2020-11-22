// Code generated by go-swagger; DO NOT EDIT.

package stock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gamersnote.com/v1/models"
)

// PutStockOKCode is the HTTP code returned for type PutStockOK
const PutStockOKCode int = 200

/*PutStockOK Successed

swagger:response putStockOK
*/
type PutStockOK struct {
}

// NewPutStockOK creates PutStockOK with default headers values
func NewPutStockOK() *PutStockOK {

	return &PutStockOK{}
}

// WriteResponse to the client
func (o *PutStockOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*PutStockDefault error

swagger:response putStockDefault
*/
type PutStockDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPutStockDefault creates PutStockDefault with default headers values
func NewPutStockDefault(code int) *PutStockDefault {
	if code <= 0 {
		code = 500
	}

	return &PutStockDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put stock default response
func (o *PutStockDefault) WithStatusCode(code int) *PutStockDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put stock default response
func (o *PutStockDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put stock default response
func (o *PutStockDefault) WithPayload(payload *models.Error) *PutStockDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put stock default response
func (o *PutStockDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutStockDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
