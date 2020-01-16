package main

import (
	"device-manager/pkg/errors/httperrors"
	"device-manager/pkg/flags"
	devices_handler "device-manager/pkg/handler/devices"
	users_handler "device-manager/pkg/handler/users"
	"device-manager/pkg/restapi/operations"
	"device-manager/pkg/restapi/operations/devices"
	"device-manager/pkg/restapi/operations/users"
	"device-manager/pkg/service"
	"device-manager/pkg/store"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"net/http"
)

func configureAPI(api *operations.DeviceManagerAPI, conf *flags.Config) http.Handler {
	st := store.NewStore()

	jwtService := service.NewJWT(conf.AppSecret)
	userService := service.NewUserService(st, jwtService)
	deviceService := service.NewDeviceService(st, jwtService)

	api.UserAuth = jwtService.ValidateUserToken
	api.DeviceAuth = jwtService.ValidateDeviceToken

	usersContext := users_handler.NewContext(userService)
	api.UsersUserRegistrationHandler = users.UserRegistrationHandlerFunc(usersContext.Register)

	deviceContext := devices_handler.NewContext(deviceService)
	api.DevicesDeviceRegistrationHandler = devices.DeviceRegistrationHandlerFunc(deviceContext.Register)
	api.DevicesDeviceListHandler = devices.DeviceListHandlerFunc(deviceContext.GetList)
	api.DevicesDeviceStatsHandler = devices.DeviceStatsHandlerFunc(deviceContext.GetStats)
	api.DevicesDeviceReadingsHandler = devices.DeviceReadingsHandlerFunc(deviceContext.SaveReadings)

	return setupMiddleware(api.Serve(setupMiddleware))
}

func setupMiddleware(handler http.Handler) http.Handler {
	return middlewareRecover(handler)
}

func middlewareRecover(h http.Handler) http.Handler {
	var applicationJSON = "application/json"
	var contentType = http.CanonicalHeaderKey("Content-Type")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {

				switch errTyped := err.(type) {
				case error:
					log.Fatal().Err(errTyped).Msg(errTyped.Error())
				case string:
					log.Fatal().Err(errors.New(errTyped)).Msg(errTyped)
				default:
					log.Fatal().Interface("err", err).Msg("panic handler")
				}

				w.WriteHeader(http.StatusInternalServerError)
				w.Header().Set(contentType, applicationJSON)

				if err := json.NewEncoder(w).Encode(httperrors.InternalServerError); err != nil {
					log.Fatal().Err(err).Msg("cant't write InternalServerError")
				}
			}
		}()
		h.ServeHTTP(w, r)

	})
}
