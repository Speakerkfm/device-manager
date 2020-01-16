package service

import (
	"device-manager/pkg/service/serviceiface"
	"device-manager/pkg/store"
)

const typeUser = "user"

type UserService struct {
	store  store.StoreInterface
	jwt serviceiface.JWT
}

func NewUserService(store store.StoreInterface, jwt serviceiface.JWT) *UserService {
	return &UserService{store: store, jwt: jwt}
}

func (us *UserService) NewUser(email string) (string, error) {
	user, err := us.store.NewUser(email)
	if err != nil {
		return "", err
	}

	return us.jwt.CreateJWT(user, typeUser)
}
