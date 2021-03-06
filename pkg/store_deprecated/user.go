package store_deprecated

import (
	"github.com/Speakerkfm/device-manager/pkg/errors/programerrors"
	"github.com/Speakerkfm/device-manager/pkg/models"
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

	user.ID = uuid.NewV4().String()
	user.Email = email

	err := st.codec.Set(&cache.Item{
		Key:        key,
		Object:     user,
		Expiration: userTTL,
	})

	return user, err
}


func (st *Store) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	key := fmt.Sprintf(userKey, email)

	if err := st.codec.Get(key, &user); err != nil {
		return nil, programerrors.NewObjectNotFound("User not found")
	}

	return user, nil
}

func (st *Store) UpdateUser(user *models.User) error {
	key := fmt.Sprintf(userKey, user.Email)

	return st.codec.Set(&cache.Item{
		Key:        key,
		Object:     user,
		Expiration: userTTL,
	})
}