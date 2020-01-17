package store

import (
	"device-manager/pkg/errors/programerrors"
	"device-manager/pkg/models"
	"github.com/satori/go.uuid"
)

func (st *Store) NewUser(email string) (*models.User, error) {
	st.mu.Lock()

	_, found := st.users[email]
	if found {
		st.mu.Unlock()

		return nil, programerrors.NewEmailTaken("Email is already taken")
	}

	user := &models.User{
		ID: uuid.NewV4().String(),
		Email: email,

		DevicesIDs: make([]string, 0),
	}

	st.users[email] = user

	st.mu.Unlock()

	return user, nil
}

func (st *Store) GetUserByEmail(email string) (*models.User, error) {
	st.mu.Lock()

	user, found := st.users[email]
	if !found {
		st.mu.Unlock()

		return nil, programerrors.NewObjectNotFound("User not found")
	}

	st.mu.Unlock()

	return user, nil
}

func (st *Store) UpdateUser(user *models.User) error {
	st.mu.Lock()

	user, found := st.users[user.Email]
	if !found {
		st.mu.Unlock()

		return programerrors.NewObjectNotFound("User not found")
	}

	st.users[user.Email] = user

	st.mu.Unlock()

	return nil
}
