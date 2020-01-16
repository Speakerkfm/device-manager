package service

import (
	"device-manager/pkg/models"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJWT_CreateJWT(t *testing.T) {
	jwt := NewJWT("secret")

	user := &models.User{ID: uuid.NewV4().String(), Email: "test@stet.moc"}

	token, err := jwt.CreateJWT(user, typeUser)
	assert.Nil(t, err)
	assert.True(t, token != "")
}