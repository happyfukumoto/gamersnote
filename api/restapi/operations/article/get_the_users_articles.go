// Code generated by go-swagger; DO NOT EDIT.

package article

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTheUsersArticlesHandlerFunc turns a function with the right signature into a get the users articles handler
type GetTheUsersArticlesHandlerFunc func(GetTheUsersArticlesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTheUsersArticlesHandlerFunc) Handle(params GetTheUsersArticlesParams) middleware.Responder {
	return fn(params)
}

// GetTheUsersArticlesHandler interface for that can handle valid get the users articles params
type GetTheUsersArticlesHandler interface {
	Handle(GetTheUsersArticlesParams) middleware.Responder
}

// NewGetTheUsersArticles creates a new http.Handler for the get the users articles operation
func NewGetTheUsersArticles(ctx *middleware.Context, handler GetTheUsersArticlesHandler) *GetTheUsersArticles {
	return &GetTheUsersArticles{Context: ctx, Handler: handler}
}

/*GetTheUsersArticles swagger:route GET /users/{username}/articles article getTheUsersArticles

GetTheUsersArticles get the users articles API

*/
type GetTheUsersArticles struct {
	Context *middleware.Context
	Handler GetTheUsersArticlesHandler
}

func (o *GetTheUsersArticles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTheUsersArticlesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
