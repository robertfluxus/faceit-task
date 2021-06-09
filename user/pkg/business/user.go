package business

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gofrs/uuid"
	usermodel "github.com/robertfluxus/faceit-task/user/pkg/domain"

	"github.com/jmoiron/sqlx"
)

var sqlTransactionOptions = &sql.TxOptions{Isolation: sql.LevelSerializable}

func (u *UserService) CreateUser(ctx context.Context, user *usermodel.User, requestId string) (*usermodel.User, error) {
	if user.ID == "" {
		uniqueID, err := uuid.NewV4()
		if err != nil {
			return nil, err
		}
		user.ID = fmt.Sprintf("%s", uniqueID)
	}

	var insertedUser *usermodel.User
	_, err := u.transactioner.InTransactionCtx(ctx, func(tx *sqlx.Tx) (result sql.Result, err error) {
		insertedUser, err = u.repository.InsertUser(ctx, user, requestId)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}, sqlTransactionOptions)
	if err != nil {
		return nil, err
	}
	return insertedUser, nil
}

func (u *UserService) ListUsers(ctx context.Context, countries []string) ([]*usermodel.User, error) {
	return nil, nil
}

func (u *UserService) UpdateUser(ctx context.Context, user *usermodel.User, updateMask []string, requestId string) (*usermodel.User, error) {
	return nil, nil
}

func (u *UserService) GetUser(ctx context.Context, userId string) (*usermodel.User, error) {
	return nil, nil
}
