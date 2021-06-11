package test

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	userpb "github.com/robertfluxus/faceit-task/user/api"
	usermodel "github.com/robertfluxus/faceit-task/user/pkg/domain"
	grpc "github.com/robertfluxus/faceit-task/user/pkg/grpc"
	mockuserservice "github.com/robertfluxus/faceit-task/user/pkg/grpc/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServiceHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUserService := mockuserservice.NewMockUserService(mockCtrl)
	serviceHandler := grpc.NewUserServiceHandler(mockUserService)
	context := context.Background()

	t.Run("TestCreateUserSuccess", func(t *testing.T) {
		externalUser := &userpb.User{
			Id:        "testuser",
			FirstName: "Robert",
			LastName:  "Marincu",
			Nickname:  "Rob_",
			Password:  "mysecretpass",
			Email:     "rob@gmail",
			Country:   "Romania",
		}
		userRequest := &userpb.CreateUserRequest{
			RequestId: "test-request",
			User:      externalUser,
		}
		expectedUser := &usermodel.User{
			ID:        "testuser",
			FirstName: "Robert",
			LastName:  "Marincu",
			Nickname:  "Rob_",
			Password:  "mysecretpass",
			Email:     "rob@gmail",
			Country:   "Romania",
		}
		mockUserService.EXPECT().
			CreateUser(context, expectedUser, "test-request").
			Return(expectedUser, nil)

		user, err := serviceHandler.CreateUser(context, userRequest)

		require.NoError(t, err)
		assert.Equal(t, externalUser, user)
	})

	t.Run("TestCreateInvalidUserFailure", func(t *testing.T) {
		externalUser := &userpb.User{
			Id:        "testuser",
			FirstName: "",
			LastName:  "Marincu",
			Nickname:  "Rob_",
			Password:  "mysecretpass",
			Email:     "rob@gmail",
			Country:   "Romania",
		}
		userRequest := &userpb.CreateUserRequest{
			RequestId: "test-request",
			User:      externalUser,
		}

		user, err := serviceHandler.CreateUser(context, userRequest)

		require.Error(t, grpc.ErrEmptyFirstName, err)
		assert.Empty(t, user)
	})
}
