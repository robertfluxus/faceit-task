package converters

import (

	"github.com/robertfluxus/faceit-task/pkg/domain"

	userpb "github.com/robertfluxus/faceit-task/user/api"
)

func ToExternalUser(internalUser *userpb.User) *userpb.User {
	return &userpb.User{
		Id: internalUser.ID,
		FirstName: internalUser.FirstName,
		LastName: internalUser.LastName,
		Nickname: internalUser.Nickname,
		Password: internalUser.Password,
		Email: internalUser.Email,
		Country: internalUser.Country,
	}
}

func ToInternalUser(externalUser *userpb.User) *user.User {
	return &user.User{
		ID: externalUser.Id,
		FirstName: externalUser.FirstName,
		LastName: externalUser.LastName,
		Nickname: externalUser.Nickname,
		Password: externalUser.Password,
		Email: externalUser.Email,
		Country: externalUser.Country,
	}	}
}

func ToExternalUsers(internalUsers []*user.User) []*userpb.User {
	protoUsers := make([]*userpb.User,len(users))
	for i, user := range users {
		protoUsers[i] = ToExternalUser(user)
	}
	return protoUsers
}