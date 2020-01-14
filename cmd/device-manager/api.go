package main

import (
	"device-manager/pkg/restapi/operations"
	"github.com/go-redis/redis"
	"net/http"
)

func configureAPI(api *operations.DeviceManagerAPI, redisClient *redis.Client) http.Handler {
	return setupMiddleware(api.Serve(setupMiddleware))
}

func setupMiddleware(handler http.Handler) http.Handler {
	return handler
}