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
	userKey = "user:%s"
	userTTL = 24 * time.Hour
)

func (st *Store) NewUser(email string) (*models.User, error) {
	user := &models.User{}
	key := fmt.Sprintf(userKey, email)

	if err := st.codec.Get(key, &user); err == nil {
		return nil, programerrors.NewEmailTaken("Email is already taken")
	}

	user.ID = uuid.NewV4()
	user.Email = email

	err := st.codec.Set(&cache.Item{
		Key:        key,
		Object:     user,
		Expiration: userTTL,
	})

	return user, err
}
