package service

import (
	"github.com/Speakerkfm/device-manager/pkg/errors/programerrors"
	"github.com/Speakerkfm/device-manager/pkg/mocks"
	"github.com/Speakerkfm/device-manager/pkg/models"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserService_NewUser(t *testing.T) {
	st := &mocks.StoreInterface{}
	jwt := &mocks.JWT{}

	us := NewUserService(st, jwt)

	//OK - new user
	userEmail := "test@mail.com"
	user := &models.User{ID: uuid.NewV4().String(), Email: userEmail}

	st.On("NewUser", userEmail).Return(user, nil)
	jwt.On("CreateJWT", user, typeUser).Return("token", nil)

	userToken, err := us.NewUser(userEmail)

	assert.Nil(t, err)
	assert.True(t, userToken != "")

	//Bad - email taken
	userEmail = "test1@mail.com"
	st.On("NewUser", userEmail).Return(nil, programerrors.NewEmailTaken("Email is already taken"))

	userToken, err = us.NewUser(userEmail)

	assert.NotNil(t, err)
}