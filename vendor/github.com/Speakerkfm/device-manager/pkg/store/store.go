package store

import (
	"device-manager/pkg/models"
	"sync"
)

type Store struct {
	mu       *sync.Mutex
	users    map[string]*models.User
	devices  map[string]*models.Device
	readings map[string][]*models.DeviceReadings
}

func NewStore() *Store {
	return &Store{
		mu: &sync.Mutex{},

		users:    make(map[string]*models.User, 10),
		devices:  make(map[string]*models.Device, 10),
		readings: make(map[string][]*models.DeviceReadings, 10),
	}
}
