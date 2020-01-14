package service

import (
	"device-manager/pkg/service/serviceiface"
	"device-manager/pkg/store"
)

type UserService struct {
	st  store.StoreInterface
	jwt serviceiface.JWT
}

func NewUserService(st store.StoreInterface, jwt serviceiface.JWT) *UserService {
	return &UserService{st: st, jwt: jwt}
}

func (us *UserService) NewUser(email string) (string, error) {
	user, err := us.st.NewUser(email)
	if err != nil {
		return "", err
	}

	return us.jwt.CreateUserJWT(user)
}
