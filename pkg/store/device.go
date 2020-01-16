package store

import (
	"device-manager/pkg/errors/programerrors"
	"device-manager/pkg/models"
	"fmt"
	"github.com/go-redis/cache"
	"github.com/satori/go.uuid"
	"time"
)

const (
	deviceKey = "device:%s"
	deviceTTL = 24 * time.Hour
)

func (st *Store) NewDevice(userID, deviceName, ownerEmail string) (*models.Device, error) {
	device := &models.Device{}

	device.ID = uuid.NewV4().String()
	device.Name = deviceName

	key := fmt.Sprintf(deviceKey, device.ID)

	err := st.codec.Set(&cache.Item{
		Key:        key,
		Object:     device,
		Expiration: deviceTTL,
	})

	return device, err
}

func (st *Store) GetDeviceByID(deviceID string) (*models.Device, error) {
	device := &models.Device{}
	key := fmt.Sprintf(deviceKey, deviceID)

	if err := st.codec.Get(key, &device); err != nil {
		return nil, programerrors.NewObjectNotFound("Device not found")
	}

	return device, nil
}

