package service

import (
	"device-manager/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	tokenExpirationTime = 24 * time.Hour
)

type JWT struct {
	secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{secret: secret}
}

func (j *JWT) CreateUserJWT(user *models.User) (string, error) {
	claims, err := structToMapClaims(user)
	if err != nil {
		return "", err
	}

	claims["type"] = typeUser

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.secret))
}

func (j *JWT) CreateDeviceJWT(user *models.User, device *models.Device) (string, error) {
	claims, err := structToMapClaims(device)
	if err != nil {
		return "", err
	}

	claims["type"] = typeDevice

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(user.Secret))
}

func (j *JWT) ValidateUserToken(token string) (*models.JWTKey, error) {
	return nil, nil
}

func (j *JWT) ValidateDeviceToken(token string) (*models.JWTKey, error) {
	return nil, nil
}
