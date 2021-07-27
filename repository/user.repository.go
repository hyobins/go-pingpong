package repository

type UserRepository interface {
	Pong() (string, error)
}
