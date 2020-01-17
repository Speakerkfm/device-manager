package service

import (
	"device-manager/pkg/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/errors"
	"reflect"
	"strings"
	"time"
)

const (
	tokenExpirationTime = 24 * time.Hour
	tokenPrefix         = "Bearer "
)

type JWT struct {
	secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{secret: secret}
}

func (j *JWT) CreateJWT(entity interface{}, entityType string) (string, error) {
	claims, err := structToMapClaims(entity)
	if err != nil {
		return "", err
	}

	claims["type"] = entityType

	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.secret))
}

func (j *JWT) ValidateUserToken(token string) (*models.AuthKey, error) {
	return j.validateToken(token, typeUser)
}

func (j *JWT) ValidateDeviceToken(token string) (*models.AuthKey, error) {
	return j.validateToken(token, typeDevice)
}

func (j *JWT) Parse(tokenString, jwtSecret string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	}
	return nil, false
}

func (j *JWT) validateToken(token, tokenType string) (*models.AuthKey, error) {
	token = strings.TrimPrefix(token, tokenPrefix)

	claims, ok := j.Parse(token, j.secret)
	if !ok {
		return nil, errors.Unauthenticated(tokenType)
	}

	if jwtType, ok := claims["type"]; ok && jwtType != tokenType {
		return nil, errors.Unauthenticated(tokenType)
	}

	deviceID := claims["id"].(string)

	entity := &models.AuthKey{ID: deviceID, Type: tokenType}

	if email, ok := claims["email"]; ok {
		userEmail := email.(string)
		entity.Email = &userEmail
	}

	return entity, nil
}

func structToMapClaims(i interface{}) (jwt.MapClaims, error) {
	now := time.Now()
	keyValues := jwt.MapClaims{
		"iat": now.Unix(),
	}

	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	typElem := reflect.TypeOf(i).Elem()

	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		structFieldName := typ.Field(i).Name
		typField, found := typElem.FieldByName(structFieldName)
		if !found {
			return keyValues, fmt.Errorf("expected field not found in jwt claims srtuct")
		}
		name := strings.Split(typField.Tag.Get("jwt"), ",")[0]
		if name == "-" {
			continue
		}

		var v interface{}
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = f.Int()
		case uint, uint8, uint16, uint32, uint64:
			v = f.Uint()
		case float32, float64:
			v = f.Float()
		case []byte:
			v = f.Bytes()
		case string:
			v = f.String()
			if v == "" {
				v = nil
			}
		}
		if v != nil {
			keyValues[name] = v
		}
	}

	return keyValues, nil
}
