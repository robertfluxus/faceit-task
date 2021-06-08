package grpc

import (
	"errors"
)

var ErrEmptyRequestID = errors.New("A request id must be provided")

var ErrEmptyUser = errors.New("A user must be provided")

var ErrEmptyFirstName = errors.New("A first name must be provided")

var ErrEmptyLastName = errors.New("A last name must be provided")

var ErrEmptyNickname = errors.New("A nickname must be provided")

var ErrEmptyPassword = errors.New("A password must be provided")

var ErrEmptyEmail = errors.New("An email must be provided")

var ErrEmptyCountry = errors.New("A country must be provided")

var ErrEmptyUserId = errors.New("A user id must be provided")

var ErrEmptyUpdateMask = errors.New("An update mask must be provided")
