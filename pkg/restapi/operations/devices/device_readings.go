// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	models "device-manager/pkg/models"
)

// DeviceReadingsHandlerFunc turns a function with the right signature into a device readings handler
type DeviceReadingsHandlerFunc func(DeviceReadingsParams, *models.JWTDeviceKey) middleware.Responder

// Handle executing the request and returning a response
func (fn DeviceReadingsHandlerFunc) Handle(params DeviceReadingsParams, principal *models.JWTDeviceKey) middleware.Responder {
	return fn(params, principal)
}

// DeviceReadingsHandler interface for that can handle valid device readings params
type DeviceReadingsHandler interface {
	Handle(DeviceReadingsParams, *models.JWTDeviceKey) middleware.Responder
}

// NewDeviceReadings creates a new http.Handler for the device readings operation
func NewDeviceReadings(ctx *middleware.Context, handler DeviceReadingsHandler) *DeviceReadings {
	return &DeviceReadings{Context: ctx, Handler: handler}
}

/*DeviceReadings swagger:route POST /devices/readings devices deviceReadings

Передача показаний устройства

*/
type DeviceReadings struct {
	Context *middleware.Context
	Handler DeviceReadingsHandler
}

func (o *DeviceReadings) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeviceReadingsParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.JWTDeviceKey
	if uprinc != nil {
		principal = uprinc.(*models.JWTDeviceKey) // this is really a models.JWTDeviceKey, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
