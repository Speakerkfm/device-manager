package service

import (
	"device-manager/pkg/errors/programerrors"
	"device-manager/pkg/models"
	"device-manager/pkg/service/serviceiface"
	"device-manager/pkg/store"
)

const typeDevice = "device"

type DeviceService struct {
	store store.StoreInterface
	jwt   serviceiface.JWT
}

func NewDeviceService(store store.StoreInterface, jwt serviceiface.JWT) *DeviceService {
	return &DeviceService{store: store, jwt: jwt}
}

func (ds *DeviceService) NewDevice(ownerEmail, deviceName string) (string, error) {
	user, err := ds.store.GetUserByEmail(ownerEmail)
	if err != nil {
		return "", err
	}

	device, err := ds.store.NewDevice(user.ID, deviceName, user.Email)
	if err != nil {
		return "", err
	}

	user.DevicesIDs = append(user.DevicesIDs, device.ID)
	if err := ds.store.UpdateUser(user); err != nil {
		return "", err
	}

	return ds.jwt.CreateJWT(device, typeDevice)
}

func (ds *DeviceService) GetUserDevices(userEmail string) ([]*models.Device, error) {
	user, err := ds.store.GetUserByEmail(userEmail)
	if err != nil {
		return nil, err
	}

	devices := make([]*models.Device, len(user.DevicesIDs))

	for idx, deviceID := range user.DevicesIDs {
		if devices[idx], err = ds.store.GetDeviceByID(deviceID); err != nil {
			return nil, programerrors.NewObjectNotFound("device not found")
		}
	}

	return devices, nil
}

func (ds *DeviceService) GetDeviceStats(deviceID string) ([]*models.DeviceReadings, error) {
	return ds.store.GetDeviceReadings(deviceID)
}

func (ds *DeviceService) SaveDeviceReadings(deviceID string, readingsTime string, temperature float64) error {
	return ds.store.AddDeviceReadings(deviceID, readingsTime, temperature)
}