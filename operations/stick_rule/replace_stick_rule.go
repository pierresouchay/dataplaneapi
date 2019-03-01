// Code generated by go-swagger; DO NOT EDIT.

package stick_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ReplaceStickRuleHandlerFunc turns a function with the right signature into a replace stick rule handler
type ReplaceStickRuleHandlerFunc func(ReplaceStickRuleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceStickRuleHandlerFunc) Handle(params ReplaceStickRuleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceStickRuleHandler interface for that can handle valid replace stick rule params
type ReplaceStickRuleHandler interface {
	Handle(ReplaceStickRuleParams, interface{}) middleware.Responder
}

// NewReplaceStickRule creates a new http.Handler for the replace stick rule operation
func NewReplaceStickRule(ctx *middleware.Context, handler ReplaceStickRuleHandler) *ReplaceStickRule {
	return &ReplaceStickRule{Context: ctx, Handler: handler}
}

/*ReplaceStickRule swagger:route PUT /services/haproxy/configuration/stick_rules/{id} StickRule replaceStickRule

Replace a Stick Rule

Replaces a Stick Rule configuration by it's ID in the specified backend.

*/
type ReplaceStickRule struct {
	Context *middleware.Context
	Handler ReplaceStickRuleHandler
}

func (o *ReplaceStickRule) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceStickRuleParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}