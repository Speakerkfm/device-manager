// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import jwt "github.com/dgrijalva/jwt-go"
import mock "github.com/stretchr/testify/mock"
import models "device-manager/pkg/models"

// JWT is an autogenerated mock type for the JWT type
type JWT struct {
	mock.Mock
}

// CreateJWT provides a mock function with given fields: entity, entityType
func (_m *JWT) CreateJWT(entity interface{}, entityType string) (string, error) {
	ret := _m.Called(entity, entityType)

	var r0 string
	if rf, ok := ret.Get(0).(func(interface{}, string) string); ok {
		r0 = rf(entity, entityType)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, string) error); ok {
		r1 = rf(entity, entityType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Parse provides a mock function with given fields: tokenString, jwtSecret
func (_m *JWT) Parse(tokenString string, jwtSecret string) (jwt.MapClaims, bool) {
	ret := _m.Called(tokenString, jwtSecret)

	var r0 jwt.MapClaims
	if rf, ok := ret.Get(0).(func(string, string) jwt.MapClaims); ok {
		r0 = rf(tokenString, jwtSecret)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(jwt.MapClaims)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string, string) bool); ok {
		r1 = rf(tokenString, jwtSecret)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// ValidateDeviceToken provides a mock function with given fields: token
func (_m *JWT) ValidateDeviceToken(token string) (*models.AuthKey, error) {
	ret := _m.Called(token)

	var r0 *models.AuthKey
	if rf, ok := ret.Get(0).(func(string) *models.AuthKey); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.AuthKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateUserToken provides a mock function with given fields: token
func (_m *JWT) ValidateUserToken(token string) (*models.AuthKey, error) {
	ret := _m.Called(token)

	var r0 *models.AuthKey
	if rf, ok := ret.Get(0).(func(string) *models.AuthKey); ok {
		r0 = rf(token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.AuthKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}