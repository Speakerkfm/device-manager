package store

import (
	"github.com/Speakerkfm/device-manager/pkg/models"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStore_AddDeviceReadings(t *testing.T) {
	st := NewStore()

	deviceID := uuid.NewV4().String()
	err := st.AddDeviceReadings(deviceID, time.Now().String(), 1.2)

	assert.Nil(t, err)
	assert.True(t, len(st.readings[deviceID]) == 1)

	err = st.AddDeviceReadings(deviceID, time.Now().String(), 1.5)

	assert.Nil(t, err)
	assert.True(t, len(st.readings[deviceID]) == 2)
}

func TestStore_GetDeviceReadings(t *testing.T) {
	st := NewStore()

	deviceID := uuid.NewV4().String()
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

	readings, err := st.GetDeviceReadings(deviceID)

	assert.Nil(t, err)
	assert.True(t, len(readings) == 3)
}
