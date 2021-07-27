package service

type UserService interface {
	PongHandler() (string, error)
}
