package business

import (
	"context"
	"database/sql"

	"github.com/robertfluxus/faceit-task/user/pkg/domain"

	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *user.User, requestID string) (*user.User, error)
	QueryUsers()
	GetUserByID()
	UpdateUser()
}

type UserService struct {
	repository    UserRepository
	transactioner Transactioner
}

func NewUserService(repository UserRepository, transactioner Transactioner) *UserService {
	return &UserService{
		repository:    repository,
		transactioner: transactioner,
	}
}

type Transactioner interface {
	WithTransaction(fn TxFn) error
}
