// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"device-manager/pkg/restapi/operations"
	"device-manager/pkg/restapi/operations/devices"
	"device-manager/pkg/restapi/operations/users"

	models "device-manager/pkg/models"
)

//go:generate swagger generate server --target ../../pkg --name DeviceManager --spec ../../tmp/swagger.yaml --tags users --tags devices --principal models.AuthKey --exclude-main

func configureFlags(api *operations.DeviceManagerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.DeviceManagerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	api.DeviceAuth = func(token string) (*models.AuthKey, error) {
		return nil, errors.NotImplemented("api key auth (Device) Authorization from header param [Authorization] has not yet been implemented")
	}
	// Applies when the "Authorization" header is set
	api.UserAuth = func(token string) (*models.AuthKey, error) {
		return nil, errors.NotImplemented("api key auth (User) Authorization from header param [Authorization] has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	api.DevicesDeviceListHandler = devices.DeviceListHandlerFunc(func(params devices.DeviceListParams, principal *models.AuthKey) middleware.Responder {
		return middleware.NotImplemented("operation devices.DeviceList has not yet been implemented")
	})
	api.DevicesDeviceReadingsHandler = devices.DeviceReadingsHandlerFunc(func(params devices.DeviceReadingsParams, principal *models.AuthKey) middleware.Responder {
		return middleware.NotImplemented("operation devices.DeviceReadings has not yet been implemented")
	})
	api.DevicesDeviceRegistrationHandler = devices.DeviceRegistrationHandlerFunc(func(params devices.DeviceRegistrationParams) middleware.Responder {
		return middleware.NotImplemented("operation devices.DeviceRegistration has not yet been implemented")
	})
	api.DevicesDeviceStatsHandler = devices.DeviceStatsHandlerFunc(func(params devices.DeviceStatsParams, principal *models.AuthKey) middleware.Responder {
		return middleware.NotImplemented("operation devices.DeviceStats has not yet been implemented")
	})
	api.UsersUserRegistrationHandler = users.UserRegistrationHandlerFunc(func(params users.UserRegistrationParams) middleware.Responder {
		return middleware.NotImplemented("operation users.UserRegistration has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
