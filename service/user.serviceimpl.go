package service

import (
	"github.com/hyobins/go-pingpong/repository"
)

type UserServiceImpl struct {
	repository.UserRepository
}

func (UserServiceImpl) NewUserServiceImpl(repo repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{repo}
}

//ping -> pong
func (userServiceImpl *UserServiceImpl) PongHandler() (string, error) {
	return userServiceImpl.UserRepository.Pong()
}
