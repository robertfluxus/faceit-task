package db

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type Connection interface {
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	sq.BaseRunner
}

type TxFn func(Connection) error

var sqBuilder = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

type DB struct {
	conn Connection
}

func New(conn Connection) *DB {
	return &DB{
		conn: conn,
	}
}

const (
	UserTableName = "users"
)

func (db *DB) WithTransaction(fn TxFn) error {
	x, err := db.Begin()
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = fn(tx)
	return err
}
