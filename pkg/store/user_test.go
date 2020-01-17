package store

import (
	"github.com/Speakerkfm/device-manager/pkg/models"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStore_NewUser(t *testing.T) {
	st := NewStore()

	userEmail := "user@mail.com"

	//OK - new user
	user, err := st.NewUser(userEmail)

	assert.Nil(t, err)
	assert.Equal(t, user.Email, userEmail)

	//Bad - email taken
	user, err = st.NewUser(userEmail)

	assert.NotNil(t, err)
}

func TestStore_GetUserByEmail(t *testing.T) {
	st := NewStore()

	userEmail := "user@mail.com"

	st.users[userEmail] = &models.User{
		ID:    uuid.NewV4().String(),
		Email: userEmail,
	}

	//OK - user found
	user, err := st.GetUserByEmail(userEmail)

	assert.Nil(t, err)
	assert.Equal(t, user.Email, userEmail)

	userEmailNotFound := "user2@mail.com"

	//Bad - user not found
	user, err = st.GetUserByEmail(userEmailNotFound)

	assert.NotNil(t, err)
}

func TestStore_UpdateUser(t *testing.T) {
	st := NewStore()

	userEmail := "user@mail.com"
	user := &models.User{
		ID:    uuid.NewV4().String(),
		Email: userEmail,
	}

	st.users[userEmail] = user

	//OK - user found
	user.DevicesIDs = []string{uuid.NewV4().String()}
	err := st.UpdateUser(user)

	assert.Nil(t, err)
	assert.True(t, len(st.users[userEmail].DevicesIDs) == 1)
}