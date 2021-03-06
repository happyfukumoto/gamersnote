// Code generated by go-swagger; DO NOT EDIT.

package article

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetArticleHandlerFunc turns a function with the right signature into a get article handler
type GetArticleHandlerFunc func(GetArticleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetArticleHandlerFunc) Handle(params GetArticleParams) middleware.Responder {
	return fn(params)
}

// GetArticleHandler interface for that can handle valid get article params
type GetArticleHandler interface {
	Handle(GetArticleParams) middleware.Responder
}

// NewGetArticle creates a new http.Handler for the get article operation
func NewGetArticle(ctx *middleware.Context, handler GetArticleHandler) *GetArticle {
	return &GetArticle{Context: ctx, Handler: handler}
}

/*GetArticle swagger:route GET /articles/{article_id} article getArticle

GetArticle get article API

*/
type GetArticle struct {
	Context *middleware.Context
	Handler GetArticleHandler
}

func (o *GetArticle) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetArticleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
