// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "device-manager/pkg/models"
)

// UserRegistrationOKCode is the HTTP code returned for type UserRegistrationOK
const UserRegistrationOKCode int = 200

/*UserRegistrationOK OK

swagger:response userRegistrationOK
*/
type UserRegistrationOK struct {

	/*
	  In: Body
	*/
	Payload *UserRegistrationOKBody `json:"body,omitempty"`
}

// NewUserRegistrationOK creates UserRegistrationOK with default headers values
func NewUserRegistrationOK() *UserRegistrationOK {

	return &UserRegistrationOK{}
}

// WithPayload adds the payload to the user registration o k response
func (o *UserRegistrationOK) WithPayload(payload *UserRegistrationOKBody) *UserRegistrationOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user registration o k response
func (o *UserRegistrationOK) SetPayload(payload *UserRegistrationOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserRegistrationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UserRegistrationUnprocessableEntityCode is the HTTP code returned for type UserRegistrationUnprocessableEntity
const UserRegistrationUnprocessableEntityCode int = 422

/*UserRegistrationUnprocessableEntity Error

swagger:response userRegistrationUnprocessableEntity
*/
type UserRegistrationUnprocessableEntity struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResult `json:"body,omitempty"`
}

// NewUserRegistrationUnprocessableEntity creates UserRegistrationUnprocessableEntity with default headers values
func NewUserRegistrationUnprocessableEntity() *UserRegistrationUnprocessableEntity {

	return &UserRegistrationUnprocessableEntity{}
}

// WithPayload adds the payload to the user registration unprocessable entity response
func (o *UserRegistrationUnprocessableEntity) WithPayload(payload *models.ErrorResult) *UserRegistrationUnprocessableEntity {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user registration unprocessable entity response
func (o *UserRegistrationUnprocessableEntity) SetPayload(payload *models.ErrorResult) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserRegistrationUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
