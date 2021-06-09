package db

import (
	"context"
	"database/sql"
	"errors"

	dbmodel "github.com/robertfluxus/faceit-task/user/pkg/db/model"
	"github.com/robertfluxus/faceit-task/user/pkg/domain"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

var userQuery = sqBuilder.Select(
	"users.id",
	"users.first_name",
	"users.last_name",
	"users.nickname",
	"users.password",
	"users.email",
	"users.country",
).From(UserTableName)

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
	var res *sql.Rows
	err := WithTransaction(db.db, func(tx Transaction) (err error) {
		res, err = sqBuilder.Insert(UserTableName).
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
			Suffix(`RETURNING id, first_name, last_name, nickname, password, email, country`).
			RunWith(db.db).
			QueryContext(ctx)
		if err != nil {
			return err
		}
		return nil
	})
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

func (db *DB) GetUserByID(ctx context.Context, userID string) (*user.User, error) {
	result, err := userQuery.
		Where(sq.Eq{"users.id": userID}).
		RunWith(db.db).
		QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var userRecords dbmodel.UserRecords
	err = sqlx.StructScan(result, &userRecords)
	if err != nil {
		return nil, err
	}
	return userRecords[0].ToDomain(), nil
}
