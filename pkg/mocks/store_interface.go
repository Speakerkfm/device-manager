// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "device-manager/pkg/models"

// StoreInterface is an autogenerated mock type for the StoreInterface type
type StoreInterface struct {
	mock.Mock
}

// AddDeviceReadings provides a mock function with given fields: deviceID, readingsTime, temperature
func (_m *StoreInterface) AddDeviceReadings(deviceID string, readingsTime string, temperature float64) error {
	ret := _m.Called(deviceID, readingsTime, temperature)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, float64) error); ok {
		r0 = rf(deviceID, readingsTime, temperature)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDeviceByID provides a mock function with given fields: deviceID
func (_m *StoreInterface) GetDeviceByID(deviceID string) (*models.Device, error) {
	ret := _m.Called(deviceID)

	var r0 *models.Device
	if rf, ok := ret.Get(0).(func(string) *models.Device); ok {
		r0 = rf(deviceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(deviceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceReadings provides a mock function with given fields: deviceID
func (_m *StoreInterface) GetDeviceReadings(deviceID string) ([]*models.DeviceReadings, error) {
	ret := _m.Called(deviceID)

	var r0 []*models.DeviceReadings
	if rf, ok := ret.Get(0).(func(string) []*models.DeviceReadings); ok {
		r0 = rf(deviceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.DeviceReadings)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(deviceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *StoreInterface) GetUserByEmail(email string) (*models.User, error) {
	ret := _m.Called(email)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(string) *models.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDevice provides a mock function with given fields: userID, deviceName, ownerEmail
func (_m *StoreInterface) NewDevice(userID string, deviceName string, ownerEmail string) (*models.Device, error) {
	ret := _m.Called(userID, deviceName, ownerEmail)

	var r0 *models.Device
	if rf, ok := ret.Get(0).(func(string, string, string) *models.Device); ok {
		r0 = rf(userID, deviceName, ownerEmail)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(userID, deviceName, ownerEmail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUser provides a mock function with given fields: email
func (_m *StoreInterface) NewUser(email string) (*models.User, error) {
	ret := _m.Called(email)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(string) *models.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: user
func (_m *StoreInterface) UpdateUser(user *models.User) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}