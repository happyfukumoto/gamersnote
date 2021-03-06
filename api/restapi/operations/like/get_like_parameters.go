// Code generated by go-swagger; DO NOT EDIT.

package like

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetLikeParams creates a new GetLikeParams object
// no default values defined in spec.
func NewGetLikeParams() GetLikeParams {

	return GetLikeParams{}
}

// GetLikeParams contains all the bound params for the get like operation
// typically these are obtained from a http.Request
//
// swagger:parameters getLike
type GetLikeParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	ArticleID string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetLikeParams() beforehand.
func (o *GetLikeParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rArticleID, rhkArticleID, _ := route.Params.GetOK("article_id")
	if err := o.bindArticleID(rArticleID, rhkArticleID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindArticleID binds and validates parameter ArticleID from path.
func (o *GetLikeParams) bindArticleID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ArticleID = raw

	return nil
}
