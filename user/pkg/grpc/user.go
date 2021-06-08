package grpc

import (
	"context"

	userpb "github.com/robertfluxus/faceit-task/user/api"
)

// CreateUser creates a new user
func (s *ServiceHandler) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (userpb.User, error) {

	err := validateCreateUserRequest(req)
	if err != nil {
		return nil, err
	}

	internalUser := converters.ToInternalUser(req.User)

	user, err := s.userService.CreateUser(ctx, user, req.RequestId)
	if err != nil {
		return nil, err
	}
	return converters.ToExternalUser(user), nil
}

func validateCreateUserRequest(req *userpb.CreateUserRequest) error {
	if req.RequestId == "" {
		return ErrEmptyRequestID
	}

	if req.User == nil {
		return ErrEmptyUser
	}

	if req.User.FirstName == "" {
		return ErrEmptyFirstName
	}

	if req.User.LastName == "" {
		return ErrEmptyLastName
	}

	if req.User.Nickname == "" {
		return ErrEmptyNickname
	}

	if req.User.Password == "" {
		return ErrEmptyPassword
	}

	if req.User.Email == "" {
		return ErrEmptyEmail
	}

	if req.User.Country == "" {
		return ErrEmptyCountry
	}

	return nil
}

func (s *ServiceHandler) GetUser(ctx context.Context, req *userpb.GetUserRequest) (userpb.User, error) {
	err := validateGetUserRequest(req)
	if err != nil {
		return nil, err
	}

	user, err := s.userService.GetUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return converters.ToExternalUser(user), nil
}

func validateGetUserRequest(req *userpb.GetUserRequest) error {
	if req.UserId == "" {
		return ErrEmptyUserId
	}
	return nil
}

func (s *ServiceHandler) ListUsers(ctx context.Context, req *userpb.ListUsersRequest) (userpb.ListUsersResponse, error) {
	err := validateListUsersRequest(req)
	if err != nil {
		return nil, err
	}

	users, err := s.userService.ListUsers(ctx, req.Country)
	if err != nil {
		return err
	}

	return &userpb.ListUsersResponse{
		Users: converters.ToExternalUsers(users),
	}, nil
}

func validateListUsersRequest(req *userpb.ListUsersRequest) error {
	if req.Country == "" {
		return ErrEmptyCountry
	}
	return nil
}

func (s *ServiceHandler) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.User, error) {
	err := validateUpdateUserRequest(req)
	if err != nil {
		return nil, err
	}

	user, err := s.userService.UpdateUser(ctx, converters.ToInternalUser(req.User), req.UpdateMask.Paths, req.RequestId)
	if err != nil {
		return nil, err
	}
	return converters.ToExternalUser(user), nil
}

func validateUpdateUserRequest(req *userpb.UpdateUserRequest) error {
	if req.RequestId == "" {
		return ErrEmptyRequestID
	}

	if req.User == nil {
		return ErrEmptyUser
	}

	if req.User.Id == "" {
		return ErrEmptyUserId
	}

	if req.UpdateMask == nil {
		return ErrEmptyUpdateMask
	}

	return nil
}
