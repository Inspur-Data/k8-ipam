// Code generated by go-swagger; DO NOT EDIT.

package ipamwrapper_agent

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteIpamHandlerFunc turns a function with the right signature into a delete ipam handler
type DeleteIpamHandlerFunc func(DeleteIpamParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteIpamHandlerFunc) Handle(params DeleteIpamParams) middleware.Responder {
	return fn(params)
}

// DeleteIpamHandler interface for that can handle valid delete ipam params
type DeleteIpamHandler interface {
	Handle(DeleteIpamParams) middleware.Responder
}

// NewDeleteIpam creates a new http.Handler for the delete ipam operation
func NewDeleteIpam(ctx *middleware.Context, handler DeleteIpamHandler) *DeleteIpam {
	return &DeleteIpam{Context: ctx, Handler: handler}
}

/*
	DeleteIpam swagger:route DELETE /ipam ipamwrapper-agent deleteIpam

# Delete ip from ipamwrapper

Send a request to ipamwrapper to delete an ip
*/
type DeleteIpam struct {
	Context *middleware.Context
	Handler DeleteIpamHandler
}

func (o *DeleteIpam) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteIpamParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
