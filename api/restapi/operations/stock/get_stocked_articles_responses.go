// Code generated by go-swagger; DO NOT EDIT.

package stock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gamersnote.com/v1/models"
)

// GetStockedArticlesOKCode is the HTTP code returned for type GetStockedArticlesOK
const GetStockedArticlesOKCode int = 200

/*GetStockedArticlesOK get stocks

swagger:response getStockedArticlesOK
*/
type GetStockedArticlesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Article `json:"body,omitempty"`
}

// NewGetStockedArticlesOK creates GetStockedArticlesOK with default headers values
func NewGetStockedArticlesOK() *GetStockedArticlesOK {

	return &GetStockedArticlesOK{}
}

// WithPayload adds the payload to the get stocked articles o k response
func (o *GetStockedArticlesOK) WithPayload(payload []*models.Article) *GetStockedArticlesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get stocked articles o k response
func (o *GetStockedArticlesOK) SetPayload(payload []*models.Article) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStockedArticlesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Article, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetStockedArticlesDefault error

swagger:response getStockedArticlesDefault
*/
type GetStockedArticlesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetStockedArticlesDefault creates GetStockedArticlesDefault with default headers values
func NewGetStockedArticlesDefault(code int) *GetStockedArticlesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetStockedArticlesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get stocked articles default response
func (o *GetStockedArticlesDefault) WithStatusCode(code int) *GetStockedArticlesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get stocked articles default response
func (o *GetStockedArticlesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get stocked articles default response
func (o *GetStockedArticlesDefault) WithPayload(payload *models.Error) *GetStockedArticlesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get stocked articles default response
func (o *GetStockedArticlesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStockedArticlesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
