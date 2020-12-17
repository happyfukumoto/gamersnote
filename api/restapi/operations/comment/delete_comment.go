// Code generated by go-swagger; DO NOT EDIT.

package comment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteCommentHandlerFunc turns a function with the right signature into a delete comment handler
type DeleteCommentHandlerFunc func(DeleteCommentParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteCommentHandlerFunc) Handle(params DeleteCommentParams) middleware.Responder {
	return fn(params)
}

// DeleteCommentHandler interface for that can handle valid delete comment params
type DeleteCommentHandler interface {
	Handle(DeleteCommentParams) middleware.Responder
}

// NewDeleteComment creates a new http.Handler for the delete comment operation
func NewDeleteComment(ctx *middleware.Context, handler DeleteCommentHandler) *DeleteComment {
	return &DeleteComment{Context: ctx, Handler: handler}
}

/*DeleteComment swagger:route DELETE /articles/{article_id}/comments/{comment_id} comment deleteComment

DeleteComment delete comment API

*/
type DeleteComment struct {
	Context *middleware.Context
	Handler DeleteCommentHandler
}

func (o *DeleteComment) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteCommentParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
