package users

import (
	"device-manager/pkg/errors/httperrors"
	"device-manager/pkg/restapi/operations/users"
	"device-manager/pkg/service/serviceiface"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rs/zerolog/log"
)

type Context struct {
	us serviceiface.UserService
}

func NewContext(us serviceiface.UserService) *Context {
	return &Context{us: us}
}

func (c *Context) Register(params users.UserRegistrationParams) middleware.Responder {
	token, err := c.us.NewUser(params.Body.Email)
	if err != nil {
		log.Warn().Err(err).Msg("Registration user failed")

		return users.NewUserRegistrationUnprocessableEntity().WithPayload(&httperrors.EmailTakenError)
	}

	return users.NewUserRegistrationOK().WithPayload(&users.UserRegistrationOKBody{Token: token})
}
