// Code generated by go-swagger; DO NOT EDIT.

package comment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PutCommentHandlerFunc turns a function with the right signature into a put comment handler
type PutCommentHandlerFunc func(PutCommentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PutCommentHandlerFunc) Handle(params PutCommentParams) middleware.Responder {
	return fn(params)
}

// PutCommentHandler interface for that can handle valid put comment params
type PutCommentHandler interface {
	Handle(PutCommentParams) middleware.Responder
}

// NewPutComment creates a new http.Handler for the put comment operation
func NewPutComment(ctx *middleware.Context, handler PutCommentHandler) *PutComment {
	return &PutComment{Context: ctx, Handler: handler}
}

/*PutComment swagger:route PUT /articles/{article_id}/comments/{comment_id} comment putComment

PutComment put comment API

*/
type PutComment struct {
	Context *middleware.Context
	Handler PutCommentHandler
}

func (o *PutComment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPutCommentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
