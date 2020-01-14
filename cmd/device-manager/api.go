package main

import (
	users_handler "device-manager/pkg/handler/users"
	"device-manager/pkg/restapi/operations"
	"device-manager/pkg/restapi/operations/users"
	"device-manager/pkg/service"
	"device-manager/pkg/store"
	"github.com/go-redis/redis"
	"net/http"
)

func configureAPI(api *operations.DeviceManagerAPI, redisClient *redis.Client) http.Handler {
	st := store.NewStore(redisClient)

	jwtService := service.NewJWT()
	userService := service.NewUserService(st, jwtService)

	usersContext := users_handler.NewContext(userService)
	api.UsersUserRegistrationHandler = users.UserRegistrationHandlerFunc(usersContext.Register)

	return setupMiddleware(api.Serve(setupMiddleware))
}

func setupMiddleware(handler http.Handler) http.Handler {
	return handler
}