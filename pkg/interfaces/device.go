// Code generated by ifacemaker. DO NOT EDIT.

package interfaces

import (
	"github.com/Speakerkfm/device-manager/pkg/models"
)

type DeviceService interface {
	NewDevice(ownerEmail, deviceName string) (string, error)
	GetUserDevices(userEmail string) ([]*models.Device, error)
	CheckUserDevice(userEmail, deviceID string) bool
	GetDeviceStats(deviceID string) ([]*models.DeviceReadings, error)
	SaveDeviceReadings(deviceID string, readingsTime string, temperature float64) error
}
