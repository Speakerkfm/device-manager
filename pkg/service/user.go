package service

import (
	"github.com/Speakerkfm/device-manager/pkg/interfaces"
)

const typeUser = "user"

type UserService struct {
	store interfaces.StoreInterface
	jwt   interfaces.JWT
}

func NewUserService(store interfaces.StoreInterface, jwt interfaces.JWT) *UserService {
	return &UserService{store: store, jwt: jwt}
}

func (us *UserService) NewUser(email string) (string, error) {
	user, err := us.store.NewUser(email)
	if err != nil {
		return "", err
	}

	return us.jwt.CreateJWT(user, typeUser)
}
