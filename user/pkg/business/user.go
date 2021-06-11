package business

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
	usermodel "github.com/robertfluxus/faceit-task/user/pkg/domain"
)

func (u *UserService) CreateUser(ctx context.Context, user *usermodel.User, requestId string) (*usermodel.User, error) {
	if user.ID == "" {
		uniqueID, err := uuid.NewV4()
		if err != nil {
			return nil, err
		}
		user.ID = fmt.Sprintf("%s", uniqueID)
	}

	var insertedUser *usermodel.User
	insertedUser, err := u.repository.InsertUser(ctx, user, requestId)
	if err != nil {
		return nil, err
	}
	return insertedUser, nil
}

func (u *UserService) ListUsers(ctx context.Context, countries []string) ([]*usermodel.User, error) {
	return nil, nil
}

func (u *UserService) UpdateUser(ctx context.Context, userUpdate *usermodel.User, updateMask []string, requestId string) (*usermodel.User, error) {
	err := validateEntityUpdate(userUpdate, updateMask)
	if err != nil {
		return nil, err
	}

	appliedUserUpdate, err := u.applyUserUpdate(ctx, userUpdate, updateMask)
	if err != nil {
		return nil, err
	}
	var updatedUser *usermodel.User
	updatedUser, err = u.repository.UpdateUser(ctx, appliedUserUpdate)
	if err != nil {
		return nil, err
	}
	err = u.postUpdateToRabbitMQ(updatedUser)
	if err != nil {
		log.Print("Failed to post message: %w", err)
	}
	return updatedUser, nil
}

func (u *UserService) postUpdateToRabbitMQ(userUpdate *usermodel.User) error {
	message, err := json.Marshal(userUpdate)
	if err != nil {
		return err
	}
	err = u.rabbit.PublishMessage(message)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) applyUserUpdate(ctx context.Context, userUpdate *usermodel.User, updateMask []string) (*usermodel.User, error) {
	existingUser, err := u.repository.GetUserByID(ctx, userUpdate.ID)
	if err != nil {
		return nil, err
	}

	for _, field := range updateMask {
		switch field {
		case usermodel.FieldCountry:
			existingUser.Country = userUpdate.Country
		case usermodel.FieldNickname:
			existingUser.Nickname = userUpdate.Nickname
		}
	}
	return existingUser, nil
}

func validateEntityUpdate(user *usermodel.User, updateMask []string) error {
	for _, field := range updateMask {
		switch field {
		case usermodel.FieldCountry:
			if user.Country == "" {
				return errors.New("supplied country update is empty")
			}
		case usermodel.FieldNickname:
			if user.Nickname == "" {
				return errors.New("supplied nickname update is empty")
			}
		default:
			return errors.New("supplied update field is unknown")
		}
	}
	return nil
}

func (u *UserService) GetUser(ctx context.Context, userID string) (*usermodel.User, error) {
	user, err := u.repository.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
