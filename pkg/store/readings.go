package store

import (
	"device-manager/pkg/models"
	"time"
)

func (st *Store) AddDeviceReadings(deviceID string, readingsTime string, temperature float64) error {
	st.mu.Lock()

	deviceReadings, found := st.readings[deviceID]
	if !found {
		deviceReadings = make([]*models.DeviceReadings, 0)
	}

	deviceReadings = append(deviceReadings,
		&models.DeviceReadings{
			MeterReadingsTime: readingsTime,
			Temperature:       temperature,
			Created:           time.Now(),
		})

	st.readings[deviceID] = deviceReadings

	st.mu.Unlock()

	return nil
}

func (st *Store) GetDeviceReadings(deviceID string) ([]*models.DeviceReadings, error) {
	st.mu.Lock()

	deviceReadings, found := st.readings[deviceID]
	if !found {
		deviceReadings = make([]*models.DeviceReadings, 0)
	}

	st.mu.Unlock()

	return deviceReadings, nil
}
