// Code generated by go-swagger; DO NOT EDIT.

package stock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gamersnote.com/v1/utils"
)

// GetStockedArticlesHandlerFunc turns a function with the right signature into a get stocked articles handler
type GetStockedArticlesHandlerFunc func(GetStockedArticlesParams, *utils.TokenPayload) middleware.Responder

// Handle executing the request and returning a response
func (fn GetStockedArticlesHandlerFunc) Handle(params GetStockedArticlesParams, principal *utils.TokenPayload) middleware.Responder {
	return fn(params, principal)
}

// GetStockedArticlesHandler interface for that can handle valid get stocked articles params
type GetStockedArticlesHandler interface {
	Handle(GetStockedArticlesParams, *utils.TokenPayload) middleware.Responder
}

// NewGetStockedArticles creates a new http.Handler for the get stocked articles operation
func NewGetStockedArticles(ctx *middleware.Context, handler GetStockedArticlesHandler) *GetStockedArticles {
	return &GetStockedArticles{Context: ctx, Handler: handler}
}

/*GetStockedArticles swagger:route GET /users/me/stocks stock getStockedArticles

GetStockedArticles get stocked articles API

*/
type GetStockedArticles struct {
	Context *middleware.Context
	Handler GetStockedArticlesHandler
}

func (o *GetStockedArticles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetStockedArticlesParams()

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
