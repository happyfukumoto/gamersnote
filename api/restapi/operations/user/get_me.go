// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gamersnote.com/v1/utils"
)

// GetMeHandlerFunc turns a function with the right signature into a get me handler
type GetMeHandlerFunc func(GetMeParams, *utils.TokenPayload) middleware.Responder

// Handle executing the request and returning a response
func (fn GetMeHandlerFunc) Handle(params GetMeParams, principal *utils.TokenPayload) middleware.Responder {
	return fn(params, principal)
}

// GetMeHandler interface for that can handle valid get me params
type GetMeHandler interface {
	Handle(GetMeParams, *utils.TokenPayload) middleware.Responder
}

// NewGetMe creates a new http.Handler for the get me operation
func NewGetMe(ctx *middleware.Context, handler GetMeHandler) *GetMe {
	return &GetMe{Context: ctx, Handler: handler}
}

/*GetMe swagger:route GET /users/me user getMe

GetMe get me API

*/
type GetMe struct {
	Context *middleware.Context
	Handler GetMeHandler
}

func (o *GetMe) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetMeParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *utils.TokenPayload
	if uprinc != nil {
		principal = uprinc.(*utils.TokenPayload) // this is really a utils.TokenPayload, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
