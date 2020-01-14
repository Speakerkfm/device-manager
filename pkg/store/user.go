package store

import (
	"device-manager/pkg/models"
	"fmt"
	"github.com/go-redis/cache"
	"github.com/satori/go.uuid"
	"time"
)

const (
	userKey = "user:%s"
	userTTL = 24 * time.Hour
)

func (st *Store) NewUser(email string) (*models.User, error) {
	user := &models.User{
		ID:    uuid.NewV4(),
		Email: email,
	}

	err := st.codec.Set(&cache.Item{
		Key:        fmt.Sprintf(userKey, user.ID),
		Object:     user,
		Expiration: userTTL,
	})

	return user, err
}
