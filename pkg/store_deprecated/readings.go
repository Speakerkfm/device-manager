package store_deprecated

import (
	"github.com/Speakerkfm/device-manager/pkg/models"
	"fmt"
	"github.com/go-redis/cache"
	"time"
)

const (
	readingsKey = "readings:%s"
	readingsTTL = 24 * time.Hour
)

func (st *Store) AddDeviceReadings(deviceID string, readingsTime string, temperature float64) error {
	var deviceReadings []*models.DeviceReadings
	key := fmt.Sprintf(readingsKey, deviceID)

	if err := st.codec.Get(key, &deviceReadings); err != nil {
		deviceReadings = make([]*models.DeviceReadings, 0)
	}

	deviceReadings = append(deviceReadings,
		&models.DeviceReadings{
			MeterReadingsTime: readingsTime,
			Temperature: temperature,
			Created: time.Now(),
	})

	return st.codec.Set(&cache.Item{
		Key:        key,
		Object:     deviceReadings,
		Expiration: readingsTTL,
	})
}

func (st *Store) GetDeviceReadings(deviceID string) ([]*models.DeviceReadings, error) {
	deviceReadings := make([]*models.DeviceReadings, 0)
	key := fmt.Sprintf(readingsKey, deviceID)

	if err := st.codec.Get(key, &deviceReadings); err != nil {
		return deviceReadings, nil
	}

	return deviceReadings, nil
}
