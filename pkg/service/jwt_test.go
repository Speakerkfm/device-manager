package service

import (
	"github.com/Speakerkfm/device-manager/pkg/models"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJWT_CreateJWT(t *testing.T) {
	jwt := NewJWT("secret")

	//OK - user JWT
	user := &models.User{ID: uuid.NewV4().String(), Email: "test@stet.moc"}

	token, err := jwt.CreateJWT(user, typeUser)
	assert.Nil(t, err)
	assert.True(t, token != "")

	//OK - device JWT
	device := &models.Device{ID: uuid.NewV4().String(), Name: "my_device"}

	token, err = jwt.CreateJWT(device, typeDevice)
	assert.Nil(t, err)
	assert.True(t, token != "")
}

func TestJWT_ValidateUserToken(t *testing.T) {
	jwt := NewJWT("secret")


	//OK - user JWT
	userToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3Q1QG1haWwucnUiLCJpYXQiOjE1NzkxODYwNzMsImlkIjoiN2E3OTg2ZTUtYTM0ZC00ZGVmLWJmZDEtNTBmOGQ0YWQzOTlmIiwidHlwZSI6InVzZXIifQ.fuJDtbRIQ7fpAxZEHp8Hn3LEUyclV8mKUrifYTcIYD4"

	user, err := jwt.ValidateUserToken(userToken)
	assert.Nil(t, err)
	assert.NotNil(t, user.Email)
	assert.True(t, *user.Email == "test5@mail.ru")
	assert.True(t, user.Type == typeUser)

	//Bad - wrong secret
	badUserToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3Q1QG1haWwucnUiLCJpYXQiOjE1NzkxODYwNzMsImlkIjoiN2E3OTg2ZTUtYTM0ZC00ZGVmLWJmZDEtNTBmOGQ0YWQzOTlmIiwidHlwZSI6InVzZXIifQ.0guuupaPOzczil6HGYe449XINj_EmtameDSwg2N65lA"

	user, err = jwt.ValidateUserToken(badUserToken)
	assert.NotNil(t, err)

	//Bad - device token
	deviceToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1NzkxODYwNzMsImlkIjoiN2E3OTg2ZTUtYTM0ZC00ZGVmLWJmZDEtNTBmOGQ0YWQzOTlmIiwidHlwZSI6ImRldmljZSJ9.fPmazw1OgANGSjGlX3tFoKLt06WlQzQH4gfjkPSfsao"

	user, err = jwt.ValidateUserToken(deviceToken)
	assert.NotNil(t, err)
}

func TestJWT_ValidateDeviceToken(t *testing.T) {
	jwt := NewJWT("secret")

	//OK - device JWT
	deviceToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1NzkxODYwNzMsImlkIjoiN2E3OTg2ZTUtYTM0ZC00ZGVmLWJmZDEtNTBmOGQ0YWQzOTlmIiwidHlwZSI6ImRldmljZSJ9.fPmazw1OgANGSjGlX3tFoKLt06WlQzQH4gfjkPSfsao"

	device, err := jwt.ValidateDeviceToken(deviceToken)
	assert.Nil(t, err)
	assert.Nil(t, device.Email)
	assert.True(t, device.ID != "")
	assert.True(t, device.Type == typeDevice)

	//Bad - wrong secret
	badDeviceToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3Q1QG1haWwucnUiLCJpYXQiOjE1NzkxODYwNzMsImlkIjoiN2E3OTg2ZTUtYTM0ZC00ZGVmLWJmZDEtNTBmOGQ0YWQzOTlmIiwidHlwZSI6InVzZXIifQ.fuJDtbRIQ7fpAxZEHp8Hn3LEUyclV8mKUrifYTcIYD4"

	device, err = jwt.ValidateDeviceToken(badDeviceToken)
	assert.NotNil(t, err)

	//Bad - user token
	userToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3Q1QG1haWwucnUiLCJpYXQiOjE1NzkxODYwNzMsImlkIjoiN2E3OTg2ZTUtYTM0ZC00ZGVmLWJmZDEtNTBmOGQ0YWQzOTlmIiwidHlwZSI6InVzZXIifQ.fuJDtbRIQ7fpAxZEHp8Hn3LEUyclV8mKUrifYTcIYD4"

	device, err = jwt.ValidateDeviceToken(userToken)
	assert.NotNil(t, err)
}