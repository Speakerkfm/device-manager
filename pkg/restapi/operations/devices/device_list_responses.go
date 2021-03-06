// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/Speakerkfm/device-manager/pkg/models"
)

// DeviceListOKCode is the HTTP code returned for type DeviceListOK
const DeviceListOKCode int = 200

/*DeviceListOK OK

swagger:response deviceListOK
*/
type DeviceListOK struct {

	/*
	  In: Body
	*/
	Payload []*DeviceListOKBodyItems0 `json:"body,omitempty"`
}

// NewDeviceListOK creates DeviceListOK with default headers values
func NewDeviceListOK() *DeviceListOK {

	return &DeviceListOK{}
}

// WithPayload adds the payload to the device list o k response
func (o *DeviceListOK) WithPayload(payload []*DeviceListOKBodyItems0) *DeviceListOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the device list o k response
func (o *DeviceListOK) SetPayload(payload []*DeviceListOKBodyItems0) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeviceListOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*DeviceListOKBodyItems0, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// DeviceListUnauthorizedCode is the HTTP code returned for type DeviceListUnauthorized
const DeviceListUnauthorizedCode int = 401

/*DeviceListUnauthorized Пользователь не аутентифицирован в системе

swagger:response deviceListUnauthorized
*/
type DeviceListUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResult `json:"body,omitempty"`
}

// NewDeviceListUnauthorized creates DeviceListUnauthorized with default headers values
func NewDeviceListUnauthorized() *DeviceListUnauthorized {

	return &DeviceListUnauthorized{}
}

// WithPayload adds the payload to the device list unauthorized response
func (o *DeviceListUnauthorized) WithPayload(payload *models.ErrorResult) *DeviceListUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the device list unauthorized response
func (o *DeviceListUnauthorized) SetPayload(payload *models.ErrorResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeviceListUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeviceListNotFoundCode is the HTTP code returned for type DeviceListNotFound
const DeviceListNotFoundCode int = 404

/*DeviceListNotFound Запрашиваемый ресурс не найден

swagger:response deviceListNotFound
*/
type DeviceListNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResult `json:"body,omitempty"`
}

// NewDeviceListNotFound creates DeviceListNotFound with default headers values
func NewDeviceListNotFound() *DeviceListNotFound {

	return &DeviceListNotFound{}
}

// WithPayload adds the payload to the device list not found response
func (o *DeviceListNotFound) WithPayload(payload *models.ErrorResult) *DeviceListNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the device list not found response
func (o *DeviceListNotFound) SetPayload(payload *models.ErrorResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeviceListNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeviceListUnprocessableEntityCode is the HTTP code returned for type DeviceListUnprocessableEntity
const DeviceListUnprocessableEntityCode int = 422

/*DeviceListUnprocessableEntity Error

swagger:response deviceListUnprocessableEntity
*/
type DeviceListUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResult `json:"body,omitempty"`
}

// NewDeviceListUnprocessableEntity creates DeviceListUnprocessableEntity with default headers values
func NewDeviceListUnprocessableEntity() *DeviceListUnprocessableEntity {

	return &DeviceListUnprocessableEntity{}
}

// WithPayload adds the payload to the device list unprocessable entity response
func (o *DeviceListUnprocessableEntity) WithPayload(payload *models.ErrorResult) *DeviceListUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the device list unprocessable entity response
func (o *DeviceListUnprocessableEntity) SetPayload(payload *models.ErrorResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeviceListUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
