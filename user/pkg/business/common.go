package business

import ()

type UserRepository interface {
	InsertUser()
	QueryUsers()
	GetUser()
	UpdateUser()
}

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}
