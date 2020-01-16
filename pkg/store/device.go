package store

import (
	"device-manager/pkg/errors/programerrors"
	"device-manager/pkg/models"
	"github.com/satori/go.uuid"
)

func (st *Store) NewDevice(userID, deviceName, ownerEmail string) (*models.Device, error) {
	st.mu.Lock()

	device := &models.Device{
		ID:   uuid.NewV4().String(),
		Name: deviceName,
	}

	st.devices[device.ID] = device

	st.mu.Unlock()

	return device, nil
}

func (st *Store) GetDeviceByID(deviceID string) (*models.Device, error) {
	st.mu.Lock()

	device, found := st.devices[deviceID]
	if !found {
		st.mu.Unlock()

		return nil, programerrors.NewObjectNotFound("Device not found")
	}

	st.mu.Unlock()

	readings, err := st.GetDeviceReadings(deviceID)
	if length := len(readings); err == nil && len(readings) > 0 {
		lastReadingsTime := readings[length - 1].Created.String()
		device.LastMeterReadingsTime = &lastReadingsTime
	}

	return device, nil
}
