// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"

	models "device-manager/pkg/models"
)

// DeviceReadingsHandlerFunc turns a function with the right signature into a device readings handler
type DeviceReadingsHandlerFunc func(DeviceReadingsParams, *models.AuthKey) middleware.Responder

// Handle executing the request and returning a response
func (fn DeviceReadingsHandlerFunc) Handle(params DeviceReadingsParams, principal *models.AuthKey) middleware.Responder {
	return fn(params, principal)
}

// DeviceReadingsHandler interface for that can handle valid device readings params
type DeviceReadingsHandler interface {
	Handle(DeviceReadingsParams, *models.AuthKey) middleware.Responder
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
	var principal *models.AuthKey
	if uprinc != nil {
		principal = uprinc.(*models.AuthKey) // this is really a models.AuthKey, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// DeviceReadingsBody device readings body
// swagger:model DeviceReadingsBody
type DeviceReadingsBody struct {

	// meter readings time
	// Required: true
	// Format: date-time
	MeterReadingsTime strfmt.DateTime `json:"meter_readings_time"`

	// temperature
	// Required: true
	Temperature float64 `json:"temperature"`
}

// Validate validates this device readings body
func (o *DeviceReadingsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMeterReadingsTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTemperature(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DeviceReadingsBody) validateMeterReadingsTime(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"meter_readings_time", "body", strfmt.DateTime(o.MeterReadingsTime)); err != nil {
		return err
	}

	if err := validate.FormatOf("body"+"."+"meter_readings_time", "body", "date-time", o.MeterReadingsTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DeviceReadingsBody) validateTemperature(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"temperature", "body", float64(o.Temperature)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *DeviceReadingsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DeviceReadingsBody) UnmarshalBinary(b []byte) error {
	var res DeviceReadingsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
