package business

import (
	"context"
	"database/sql"

	"github.com/robertfluxus/faceit-task/user/pkg/domain"
)

type UserRepository interface {
	InsertUser(ctx context.Context, user *user.User, requestID string) (*user.User, error)
	QueryUsers()
	GetUserByID(ctx context.Context, userID string) (*user.User, error)
	UpdateUser(ctx context.Context, user *user.User) (updatedUser *user.User, err error)
}

type UserService struct {
	repository UserRepository
	db         *sql.DB
	rabbit     *RabbitMQ
}

func NewUserService(repository UserRepository, db *sql.DB, rabbit *RabbitMQ) *UserService {
	return &UserService{
		repository: repository,
		db:         db,
		rabbit:     rabbit,
	}
}
