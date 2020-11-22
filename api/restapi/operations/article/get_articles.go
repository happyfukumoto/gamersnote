// Code generated by go-swagger; DO NOT EDIT.

package article

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetArticlesHandlerFunc turns a function with the right signature into a get articles handler
type GetArticlesHandlerFunc func(GetArticlesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetArticlesHandlerFunc) Handle(params GetArticlesParams) middleware.Responder {
	return fn(params)
}

// GetArticlesHandler interface for that can handle valid get articles params
type GetArticlesHandler interface {
	Handle(GetArticlesParams) middleware.Responder
}

// NewGetArticles creates a new http.Handler for the get articles operation
func NewGetArticles(ctx *middleware.Context, handler GetArticlesHandler) *GetArticles {
	return &GetArticles{Context: ctx, Handler: handler}
}

/*GetArticles swagger:route GET /articles article getArticles

GetArticles get articles API

*/
type GetArticles struct {
	Context *middleware.Context
	Handler GetArticlesHandler
}

func (o *GetArticles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetArticlesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
