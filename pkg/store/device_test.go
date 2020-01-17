package store

import (
	"device-manager/pkg/models"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStore_NewDevice(t *testing.T) {
	st := NewStore()

	deviceName := "my_device"
	userEmail := "user@mail.ru"
	device, err := st.NewDevice(uuid.NewV4().String(), deviceName, userEmail)

	assert.Nil(t, err)
	assert.Equal(t, device.Name, deviceName)
}

func TestStore_GetDeviceByID(t *testing.T) {
	st := NewStore()

	deviceID := uuid.NewV4().String()
	deviceName := "my_device"

	st.devices[deviceID] = &models.Device{ID: deviceID, Name: deviceName}
	st.readings[deviceID] = []*models.DeviceReadings{
		{
			MeterReadingsTime: time.Now().String(),
			Temperature:       15.5,
			Created:           time.Now(),
		},
		{
			MeterReadingsTime: time.Now().String(),
			Temperature:       15.4,
			Created:           time.Now(),
		},
		{
			MeterReadingsTime: time.Now().String(),
			Temperature:       15.4,
			Created:           time.Now(),
		},
	}

	device, err := st.GetDeviceByID(deviceID)

	assert.Nil(t, err)
	assert.Equal(t, device.Name, deviceName)
}