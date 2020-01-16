package store

import (
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

func (st *Store) NewDevice(user *models.User, deviceName, ownerEmail string) (*models.Device, error) {
	device := &models.Device{}

	device.ID = uuid.NewV4().String()
	device.Name = deviceName
	device.OwnerID = user.ID

	key := fmt.Sprintf(deviceKey, device.ID)

	err := st.codec.Set(&cache.Item{
		Key:        key,
		Object:     device,
		Expiration: deviceTTL,
	})

	return device, err
}
