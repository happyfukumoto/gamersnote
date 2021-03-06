// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PatchUserSignoutedHandlerFunc turns a function with the right signature into a patch user signouted handler
type PatchUserSignoutedHandlerFunc func(PatchUserSignoutedParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PatchUserSignoutedHandlerFunc) Handle(params PatchUserSignoutedParams) middleware.Responder {
	return fn(params)
}

// PatchUserSignoutedHandler interface for that can handle valid patch user signouted params
type PatchUserSignoutedHandler interface {
	Handle(PatchUserSignoutedParams) middleware.Responder
}

// NewPatchUserSignouted creates a new http.Handler for the patch user signouted operation
func NewPatchUserSignouted(ctx *middleware.Context, handler PatchUserSignoutedHandler) *PatchUserSignouted {
	return &PatchUserSignouted{Context: ctx, Handler: handler}
}

/*PatchUserSignouted swagger:route PATCH /users/me/signouted user patchUserSignouted

PatchUserSignouted patch user signouted API

*/
type PatchUserSignouted struct {
	Context *middleware.Context
	Handler PatchUserSignoutedHandler
}

func (o *PatchUserSignouted) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPatchUserSignoutedParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
