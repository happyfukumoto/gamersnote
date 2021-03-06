// Code generated by go-swagger; DO NOT EDIT.

package article

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostArticleHandlerFunc turns a function with the right signature into a post article handler
type PostArticleHandlerFunc func(PostArticleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostArticleHandlerFunc) Handle(params PostArticleParams) middleware.Responder {
	return fn(params)
}

// PostArticleHandler interface for that can handle valid post article params
type PostArticleHandler interface {
	Handle(PostArticleParams) middleware.Responder
}

// NewPostArticle creates a new http.Handler for the post article operation
func NewPostArticle(ctx *middleware.Context, handler PostArticleHandler) *PostArticle {
	return &PostArticle{Context: ctx, Handler: handler}
}

/*PostArticle swagger:route POST /articles article postArticle

PostArticle post article API

*/
type PostArticle struct {
	Context *middleware.Context
	Handler PostArticleHandler
}

func (o *PostArticle) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostArticleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
