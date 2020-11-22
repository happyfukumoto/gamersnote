// Code generated by go-swagger; DO NOT EDIT.

package like

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"gamersnote.com/v1/utils"
)

// GetLikedArticlesHandlerFunc turns a function with the right signature into a get liked articles handler
type GetLikedArticlesHandlerFunc func(GetLikedArticlesParams, *utils.TokenPayload) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLikedArticlesHandlerFunc) Handle(params GetLikedArticlesParams, principal *utils.TokenPayload) middleware.Responder {
	return fn(params, principal)
}

// GetLikedArticlesHandler interface for that can handle valid get liked articles params
type GetLikedArticlesHandler interface {
	Handle(GetLikedArticlesParams, *utils.TokenPayload) middleware.Responder
}

// NewGetLikedArticles creates a new http.Handler for the get liked articles operation
func NewGetLikedArticles(ctx *middleware.Context, handler GetLikedArticlesHandler) *GetLikedArticles {
	return &GetLikedArticles{Context: ctx, Handler: handler}
}

/*GetLikedArticles swagger:route GET /users/me/likes like getLikedArticles

GetLikedArticles get liked articles API

*/
type GetLikedArticles struct {
	Context *middleware.Context
	Handler GetLikedArticlesHandler
}

func (o *GetLikedArticles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetLikedArticlesParams()

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
