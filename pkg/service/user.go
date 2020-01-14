package service

import "device-manager/pkg/store"

type UserService struct {
	st store.StoreInterface
}

func NewUserService(st store.StoreInterface) *UserService {
	return &UserService{st: st}
}

func (us *UserService) NewUser(email string) (string, error) {

}
