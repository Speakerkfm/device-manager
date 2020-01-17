package devices

import (
	"github.com/Speakerkfm/device-manager/pkg/errors/httperrors"
	"github.com/Speakerkfm/device-manager/pkg/errors/programerrors"
	"github.com/Speakerkfm/device-manager/pkg/interfaces"
	"github.com/Speakerkfm/device-manager/pkg/models"
	"github.com/Speakerkfm/device-manager/pkg/restapi/operations/devices"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type Context struct {
	deviceService interfaces.DeviceService
}

func NewContext(deviceService interfaces.DeviceService) *Context {
	return &Context{deviceService: deviceService}
}

func (c *Context) Register(params devices.DeviceRegistrationParams) middleware.Responder {
	token, err := c.deviceService.NewDevice(params.Body.OwnerEmail.String(), params.Body.DeviceName)
	if err != nil {
		return getErrorResponse(err)
	}

	return devices.NewDeviceRegistrationOK().WithPayload(&devices.DeviceRegistrationOKBody{Token: token})
}

func (c *Context) GetList(params devices.DeviceListParams, user *models.AuthKey) middleware.Responder {
	userDevices, err := c.deviceService.GetUserDevices(*user.Email)
	if err != nil {
		return getErrorResponse(err)
	}

	devicesList := make([]*devices.DeviceListOKBodyItems0, len(userDevices))
	for idx, userDevice := range userDevices {
		devicesList[idx] = &devices.DeviceListOKBodyItems0{
			DeviceID:              userDevice.ID,
			DeviceName:            userDevice.Name,
			LastMeterReadingsTime: userDevice.LastMeterReadingsTime,
		}
	}

	return devices.NewDeviceListOK().WithPayload(devicesList)
}

func (c *Context) GetStats(params devices.DeviceStatsParams, user *models.AuthKey) middleware.Responder {
	if !c.deviceService.CheckUserDevice(*user.Email, params.DeviceID.String()) {
		return devices.NewDeviceStatsUnauthorized().WithPayload(&httperrors.UnauthorizedError)
	}

	deviceReadings, err := c.deviceService.GetDeviceStats(params.DeviceID.String())
	if err != nil {
		getErrorResponse(err)
	}

	return devices.NewDeviceStatsOK().WithPayload(deviceReadings)
}

func (c *Context) SaveReadings(params devices.DeviceReadingsParams, device *models.AuthKey) middleware.Responder {
	if err := c.deviceService.SaveDeviceReadings(device.ID, params.Body.MeterReadingsTime.String(), params.Body.Temperature); err != nil {
		return getErrorResponse(err)
	}

	return devices.NewDeviceReadingsCreated()
}

func getErrorResponse(err error) middleware.Responder {
	eventWarn := log.Warn()
	eventError := log.Error()

	switch errors.Cause(err).(type) {
	case programerrors.ObjectNotFound:
		eventWarn.Err(err).Msg("Registration device failed")

		return devices.NewDeviceRegistrationNotFound().WithPayload(&httperrors.ObjectNotFound)
	default:
		eventError.Err(err).Msg("Registration device failed")

		return devices.NewDeviceRegistrationUnprocessableEntity()
	}
}
