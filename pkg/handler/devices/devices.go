package devices

import (
	"device-manager/pkg/errors/httperrors"
	"device-manager/pkg/errors/programerrors"
	"device-manager/pkg/restapi/operations/devices"
	"device-manager/pkg/service/serviceiface"
	"device-manager/pkg/store"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type Context struct {
	deviceService    serviceiface.DeviceService
	store store.StoreInterface
}

func NewContext(deviceService serviceiface.DeviceService, store store.StoreInterface) *Context {
	return &Context{deviceService: deviceService, store: store}
}

func (c *Context) Register(params devices.DeviceRegistrationParams) middleware.Responder {
	user, err := c.store.GetUserByEmail(params.Body.OwnerEmail.String())
	if err != nil {
		return getErrorResponse(err)
	}

	token, err := c.deviceService.NewDevice(user, params.Body.DeviceName)
	if err != nil {
		return getErrorResponse(err)
	}

	return devices.NewDeviceRegistrationOK().WithPayload(&devices.DeviceRegistrationOKBody{Token: token})
}

func getErrorResponse(err error) middleware.Responder {
	warn := log.Warn()

	switch errors.Cause(err).(type) {
	case programerrors.ObjectNotFount:
		warn.Err(err).Msg("Registration device failed")

		return devices.NewDeviceRegistrationNotFound().WithPayload(&httperrors.ObjectNotFound)
	default:
		log.Error().Err(err).Msg("Registration device failed")

		return devices.NewDeviceRegistrationUnprocessableEntity()
	}
}