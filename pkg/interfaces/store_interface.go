// Code generated by ifacemaker. DO NOT EDIT.

package interfaces

import (
	"github.com/Speakerkfm/device-manager/pkg/models"
)

type StoreInterface interface {
	NewDevice(userID, deviceName, ownerEmail string) (*models.Device, error)
	GetDeviceByID(deviceID string) (*models.Device, error)
	AddDeviceReadings(deviceID string, readingsTime string, temperature float64) error
	GetDeviceReadings(deviceID string) ([]*models.DeviceReadings, error)
	NewUser(email string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(user *models.User) error
}
