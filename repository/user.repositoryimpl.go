package repository

type UserRepositoryImpl struct {
	data string
}

func (UserRepositoryImpl) NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (userRepositoryImpl *UserRepositoryImpl) Pong() (string, error) {
	result := "pong!"
	return result, nil
}
