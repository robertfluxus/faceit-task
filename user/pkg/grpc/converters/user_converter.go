package converters

import (
	"github.com/robertfluxus/faceit-task/user/pkg/domain"

	userpb "github.com/robertfluxus/faceit-task/user/api"
)

func ToExternalUser(internalUser *user.User) *userpb.User {
	return &userpb.User{
		Id:        internalUser.ID,
		FirstName: internalUser.FirstName,
		LastName:  internalUser.LastName,
		Nickname:  internalUser.Nickname,
		Password:  internalUser.Password,
		Email:     internalUser.Email,
		Country:   internalUser.Country,
	}
}

func ToInternalUser(externalUser *userpb.User) *user.User {
	return &user.User{
		ID:        externalUser.Id,
		FirstName: externalUser.FirstName,
		LastName:  externalUser.LastName,
		Nickname:  externalUser.Nickname,
		Password:  externalUser.Password,
		Email:     externalUser.Email,
		Country:   externalUser.Country,
	}
}

func ToExternalUsers(internalUsers []*user.User) []*userpb.User {
	protoUsers := make([]*userpb.User, len(internalUsers))
	for i, user := range internalUsers {
		protoUsers[i] = ToExternalUser(user)
	}
	return protoUsers
}
