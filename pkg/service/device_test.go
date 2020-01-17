package service

import (
	"github.com/Speakerkfm/device-manager/pkg/errors/programerrors"
	"github.com/Speakerkfm/device-manager/pkg/mocks"
	"github.com/Speakerkfm/device-manager/pkg/models"
	"github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDeviceService_NewDevice(t *testing.T) {
	jwt := &mocks.JWT{}
	st := &mocks.StoreInterface{}

	ds := NewDeviceService(st, jwt)

	//OK - new device
	userEmail := "test@mail.ru"
	deviceName := "my_device"
	user := &models.User{ID: uuid.NewV4().String(), Email: userEmail}
	device := &models.Device{ID: uuid.NewV4().String(), Name: deviceName}

	st.On("GetUserByEmail", userEmail).Return(user, nil)
	st.On("NewDevice", user.ID, deviceName, user.Email).Return(device, nil)

	user.DevicesIDs = []string{device.ID}
	st.On("UpdateUser", user).Return(nil)

	jwt.On("CreateJWT", device, typeDevice).Return("token", nil)

	deviceToken, err := ds.NewDevice(userEmail, deviceName)

	assert.Nil(t, err)
	assert.True(t, deviceToken != "")

	//Bad - user not found
	userEmail = "test1@mail.ru"
	deviceName = "my_device"

	st.On("GetUserByEmail", userEmail).Return(nil, programerrors.NewObjectNotFound("User not found"))

	deviceToken, err = ds.NewDevice(userEmail, deviceName)

	assert.NotNil(t, err)
}

func TestDeviceService_GetUserDevices(t *testing.T) {
	jwt := &mocks.JWT{}
	st := &mocks.StoreInterface{}

	ds := NewDeviceService(st, jwt)

	//OK - get devices
	userEmail := "test@mail.ru"

	user := &models.User{ID: uuid.NewV4().String(), Email: userEmail}
	device1 := &models.Device{ID: uuid.NewV4().String(), Name: "device1"}
	device2 := &models.Device{ID: uuid.NewV4().String(), Name: "device2"}

	user.DevicesIDs = []string{device1.ID, device2.ID}

	st.On("GetUserByEmail", userEmail).Return(user, nil)
	st.On("GetDeviceByID", device1.ID).Return(device1, nil)
	st.On("GetDeviceByID", device2.ID).Return(device2, nil)

	devices, err := ds.GetUserDevices(userEmail)

	assert.Nil(t, err)
	assert.True(t, len(devices) == 2)
}

func TestDeviceService_CheckUserDevice(t *testing.T) {
	jwt := &mocks.JWT{}
	st := &mocks.StoreInterface{}

	ds := NewDeviceService(st, jwt)

	//OK - check device
	userEmail := "test@mail.ru"

	user := &models.User{ID: uuid.NewV4().String(), Email: userEmail}
	device1 := &models.Device{ID: uuid.NewV4().String(), Name: "device1"}

	user.DevicesIDs = []string{device1.ID}

	st.On("GetUserByEmail", userEmail).Return(user, nil)

	userHasDevice := ds.CheckUserDevice(userEmail, device1.ID)

	assert.True(t, userHasDevice)

	//Bad - not user's device
	userEmail = "test@mail.ru"

	user = &models.User{ID: uuid.NewV4().String(), Email: userEmail}
	device1 = &models.Device{ID: uuid.NewV4().String(), Name: "device1"}

	user.DevicesIDs = []string{}

	st.On("GetUserByEmail", userEmail).Return(user, nil)

	userHasDevice = ds.CheckUserDevice(userEmail, device1.ID)

	assert.False(t, userHasDevice)
}

func TestDeviceService_GetDeviceStats(t *testing.T) {
	jwt := &mocks.JWT{}
	st := &mocks.StoreInterface{}

	ds := NewDeviceService(st, jwt)

	//OK
	deviceID := uuid.NewV4().String()

	deviceReadings1 := &models.DeviceReadings{MeterReadingsTime: time.Now().String(), Temperature: 1.2, Created: time.Now()}
	deviceReadings2 := &models.DeviceReadings{MeterReadingsTime: time.Now().String(), Temperature: 1.5, Created: time.Now()}

	st.On("GetDeviceReadings", deviceID).Return([]*models.DeviceReadings{deviceReadings1, deviceReadings2}, nil)

	deviceStats, err := ds.GetDeviceStats(deviceID)

	assert.Nil(t, err)
	assert.True(t, len(deviceStats) == 2)
}

func TestDeviceService_SaveDeviceReadings(t *testing.T) {
	jwt := &mocks.JWT{}
	st := &mocks.StoreInterface{}

	ds := NewDeviceService(st, jwt)

	//OK
	deviceID := uuid.NewV4().String()
	now := time.Now()

	st.On("AddDeviceReadings", deviceID, now.String(), 1.2).Return(nil)

	err := ds.SaveDeviceReadings(deviceID, now.String(), 1.2)

	assert.Nil(t, err)
}