package service

import (
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

func (ds *DeviceService) NewDevice(user *models.User, deviceName string) (string, error) {
	device, err := ds.store.NewDevice(user, deviceName, user.Email)
	if err != nil {
		return "", err
	}

	return ds.jwt.CreateDeviceJWT(user, device)
}
