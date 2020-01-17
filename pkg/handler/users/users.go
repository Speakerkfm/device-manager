package users

import (
	"github.com/Speakerkfm/device-manager/pkg/errors/httperrors"
	"github.com/Speakerkfm/device-manager/pkg/errors/programerrors"
	"github.com/Speakerkfm/device-manager/pkg/interfaces"
	"github.com/Speakerkfm/device-manager/pkg/restapi/operations/users"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type Context struct {
	userService interfaces.UserService
}

func NewContext(userService interfaces.UserService) *Context {
	return &Context{userService: userService}
}

func (c *Context) Register(params users.UserRegistrationParams) middleware.Responder {
	token, err := c.userService.NewUser(params.Body.Email.String())
	if err != nil {
		return getErrorResponse(err)
	}

	return users.NewUserRegistrationOK().WithPayload(&users.UserRegistrationOKBody{Token: token})
}

func getErrorResponse(err error) middleware.Responder {
	warn := log.Warn()

	switch errors.Cause(err).(type) {
	case programerrors.EmailTaken:
		warn.Err(err).Msg("Registration user failed")

		return users.NewUserRegistrationUnprocessableEntity().WithPayload(&httperrors.EmailTakenError)
	default:
		log.Error().Err(err).Msg("Registration user failed")

		return users.NewUserRegistrationUnprocessableEntity()
	}
}
