// Code generated by go-swagger; DO NOT EDIT.

package like

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gamersnote.com/v1/models"
)

// GetLikeOKCode is the HTTP code returned for type GetLikeOK
const GetLikeOKCode int = 200

/*GetLikeOK Found

swagger:response getLikeOK
*/
type GetLikeOK struct {
}

// NewGetLikeOK creates GetLikeOK with default headers values
func NewGetLikeOK() *GetLikeOK {

	return &GetLikeOK{}
}

// WriteResponse to the client
func (o *GetLikeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetLikeUnauthorizedCode is the HTTP code returned for type GetLikeUnauthorized
const GetLikeUnauthorizedCode int = 401

/*GetLikeUnauthorized Unauthorized

swagger:response getLikeUnauthorized
*/
type GetLikeUnauthorized struct {
}

// NewGetLikeUnauthorized creates GetLikeUnauthorized with default headers values
func NewGetLikeUnauthorized() *GetLikeUnauthorized {

	return &GetLikeUnauthorized{}
}

// WriteResponse to the client
func (o *GetLikeUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetLikeNotFoundCode is the HTTP code returned for type GetLikeNotFound
const GetLikeNotFoundCode int = 404

/*GetLikeNotFound NotFound

swagger:response getLikeNotFound
*/
type GetLikeNotFound struct {
}

// NewGetLikeNotFound creates GetLikeNotFound with default headers values
func NewGetLikeNotFound() *GetLikeNotFound {

	return &GetLikeNotFound{}
}

// WriteResponse to the client
func (o *GetLikeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

/*GetLikeDefault error

swagger:response getLikeDefault
*/
type GetLikeDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetLikeDefault creates GetLikeDefault with default headers values
func NewGetLikeDefault(code int) *GetLikeDefault {
	if code <= 0 {
		code = 500
	}

	return &GetLikeDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get like default response
func (o *GetLikeDefault) WithStatusCode(code int) *GetLikeDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get like default response
func (o *GetLikeDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get like default response
func (o *GetLikeDefault) WithPayload(payload *models.Error) *GetLikeDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get like default response
func (o *GetLikeDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetLikeDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
