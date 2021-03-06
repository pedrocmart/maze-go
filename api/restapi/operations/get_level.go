// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetLevelHandlerFunc turns a function with the right signature into a get level handler
type GetLevelHandlerFunc func(GetLevelParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLevelHandlerFunc) Handle(params GetLevelParams) middleware.Responder {
	return fn(params)
}

// GetLevelHandler interface for that can handle valid get level params
type GetLevelHandler interface {
	Handle(GetLevelParams) middleware.Responder
}

// NewGetLevel creates a new http.Handler for the get level operation
func NewGetLevel(ctx *middleware.Context, handler GetLevelHandler) *GetLevel {
	return &GetLevel{Context: ctx, Handler: handler}
}

/* GetLevel swagger:route GET /level getLevel

Get all levels

*/
type GetLevel struct {
	Context *middleware.Context
	Handler GetLevelHandler
}

func (o *GetLevel) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetLevelParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
