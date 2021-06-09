package grpc

import (
	"context"

	userpb "github.com/robertfluxus/faceit-task/user/api"
	"github.com/robertfluxus/faceit-task/user/pkg/domain"
)

type UserService interface {
	CreateUser(ctx context.Context, user *user.User, requestId string) (*user.User, error)
	ListUsers(ctx context.Context, countries []string) ([]*user.User, error)
	UpdateUser(ctx context.Context, user *user.User, updateMask []string, requestId string) (*user.User, error)
	GetUser(ctx context.Context, userId string) (*user.User, error)
}

type ServiceHandler struct {
	userpb.UnimplementedUserServiceServer
	userService UserService
}

func NewUserServiceHandler(userService UserService) userpb.UserServiceServer {
	return &ServiceHandler{
		userService: userService,
	}
}
