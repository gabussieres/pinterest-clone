// Code generated by go-swagger; DO NOT EDIT.

package pinterest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// FeedHandlerFunc turns a function with the right signature into a feed handler
type FeedHandlerFunc func(FeedParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FeedHandlerFunc) Handle(params FeedParams) middleware.Responder {
	return fn(params)
}

// FeedHandler interface for that can handle valid feed params
type FeedHandler interface {
	Handle(FeedParams) middleware.Responder
}

// NewFeed creates a new http.Handler for the feed operation
func NewFeed(ctx *middleware.Context, handler FeedHandler) *Feed {
	return &Feed{Context: ctx, Handler: handler}
}

/*Feed swagger:route GET /feed pinterest feed

Feed feed API

*/
type Feed struct {
	Context *middleware.Context
	Handler FeedHandler
}

func (o *Feed) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFeedParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}