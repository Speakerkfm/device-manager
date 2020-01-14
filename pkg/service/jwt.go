package service

import (
	"device-manager/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	tokenExpirationTime = 24 * 60 * 60 * time.Second
)

type JWT struct {
	secret string
}

func NewJWT() *JWT {
	return &JWT{secret: "asdfqwerqe"}
}

func (j *JWT) CreateUserJWT(user *models.User) (string, error){
	claims, err := structToMapClaims(user)
	if err != nil {
		return "", err
	}

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.secret))
}