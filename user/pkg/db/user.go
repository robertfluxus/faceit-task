package db

import (
	"context"
	"errors"

	dbmodel "github.com/robertfluxus/faceit-task/user/pkg/db/model"
	"github.com/robertfluxus/faceit-task/user/pkg/domain"

	"github.com/jmoiron/sqlx"
)

func (db *DB) InsertUser(ctx context.Context, user *user.User, requestID string) (*user.User, error) {
	userRecords, err := db.insertUserRecord(ctx, dbmodel.UserRecordFromDomain(user), requestID)
	if err != nil {
		return nil, err
	}

	if len(userRecords) == 0 {
		return nil, errors.New("Error inserting user")
	}
	return userRecords[0].ToDomain(), nil
}

func (db *DB) insertUserRecord(ctx context.Context, user *dbmodel.UserRecord, requestID string) ([]*dbmodel.UserRecord, error) {
	res, err := sqBuilder.Insert(UserTableName).
		Columns("id", "request_id", "first_name", "last_name", "nickname", "password", "email", "country").
		Values(
			user.ID,
			requestID,
			user.FirstName,
			user.LastName,
			user.Nickname,
			user.Password,
			user.Email,
			user.Country,
		).
		Suffix(`RETURNIG id, first_name, last_name, nickname, password, email, country`).
		RunWith(db.conn).
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var userRecords dbmodel.UserRecords
	err = sqlx.StructScan(res, &userRecords)
	if err != nil {
		return nil, err
	}
	return userRecords, nil
}

func (db *DB) QueryUsers() {
	return
}

func (db *DB) UpdateUser() {
	return
}

func (db *DB) GetUserByID() {
	return
}
