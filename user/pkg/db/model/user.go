package user

import (
	"github.com/robertfluxus/faceit-task/user/pkg/domain"
)

type UserRecord struct {
	ID        string `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Nickname  string `db:"nickname"`
	Password  string `db:"password"`
	Email     string `db:"email"`
	Country   string `db:"country"`
}

type UserRecords []*UserRecord

func (userRecord *UserRecord) ToDomain() *user.User {
	return &user.User{
		ID:        userRecord.ID,
		FirstName: userRecord.FirstName,
		LastName:  userRecord.LastName,
		Nickname:  userRecord.Nickname,
		Password:  userRecord.Password,
		Email:     userRecord.Email,
		Country:   userRecord.Country,
	}
}

func (userRecords UserRecords) ToDomain() []*user.User {
	var domainUsers []*user.User
	for _, r := range userRecords {
		domainUsers = append(domainUsers, r.ToDomain())
	}
	return domainUsers
}

func UserRecordFromDomain(user *user.User) *UserRecord {
	return &UserRecord{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Nickname:  user.Nickname,
		Password:  user.Password,
		Email:     user.Email,
		Country:   user.Country,
	}
}
