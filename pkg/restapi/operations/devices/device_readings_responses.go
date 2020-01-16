// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "device-manager/pkg/models"
)

// DeviceReadingsCreatedCode is the HTTP code returned for type DeviceReadingsCreated
const DeviceReadingsCreatedCode int = 201

/*DeviceReadingsCreated No content

swagger:response deviceReadingsCreated
*/
type DeviceReadingsCreated struct {
}

// NewDeviceReadingsCreated creates DeviceReadingsCreated with default headers values
func NewDeviceReadingsCreated() *DeviceReadingsCreated {

	return &DeviceReadingsCreated{}
}

// WriteResponse to the client
func (o *DeviceReadingsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// DeviceReadingsUnauthorizedCode is the HTTP code returned for type DeviceReadingsUnauthorized
const DeviceReadingsUnauthorizedCode int = 401

/*DeviceReadingsUnauthorized Пользователь не аутентифицирован в системе

swagger:response deviceReadingsUnauthorized
*/
type DeviceReadingsUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResult `json:"body,omitempty"`
}

// NewDeviceReadingsUnauthorized creates DeviceReadingsUnauthorized with default headers values
func NewDeviceReadingsUnauthorized() *DeviceReadingsUnauthorized {

	return &DeviceReadingsUnauthorized{}
}

// WithPayload adds the payload to the device readings unauthorized response
func (o *DeviceReadingsUnauthorized) WithPayload(payload *models.ErrorResult) *DeviceReadingsUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the device readings unauthorized response
func (o *DeviceReadingsUnauthorized) SetPayload(payload *models.ErrorResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeviceReadingsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeviceReadingsUnprocessableEntityCode is the HTTP code returned for type DeviceReadingsUnprocessableEntity
const DeviceReadingsUnprocessableEntityCode int = 422

/*DeviceReadingsUnprocessableEntity Error

swagger:response deviceReadingsUnprocessableEntity
*/
type DeviceReadingsUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResult `json:"body,omitempty"`
}

// NewDeviceReadingsUnprocessableEntity creates DeviceReadingsUnprocessableEntity with default headers values
func NewDeviceReadingsUnprocessableEntity() *DeviceReadingsUnprocessableEntity {

	return &DeviceReadingsUnprocessableEntity{}
}

// WithPayload adds the payload to the device readings unprocessable entity response
func (o *DeviceReadingsUnprocessableEntity) WithPayload(payload *models.ErrorResult) *DeviceReadingsUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the device readings unprocessable entity response
func (o *DeviceReadingsUnprocessableEntity) SetPayload(payload *models.ErrorResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeviceReadingsUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}