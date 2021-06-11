package common

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	amqp "github.com/streadway/amqp"
)

type ConnectionOptions struct {
	Retries       int
	Delay         time.Duration
	Timeout       time.Duration
	ConnectionURL string
}

func NewDefaultConnectionOptions(connURL string) ConnectionOptions {
	return ConnectionOptions{
		Retries:       10,
		Delay:         5 * time.Second,
		Timeout:       5 * time.Minute,
		ConnectionURL: connURL,
	}
}

func ConnectPostgress(ctx context.Context, options ConnectionOptions) (db *sql.DB, err error) {
	connErrCh := make(chan error, 1)
	defer close(connErrCh)
	go func() {
		try := 0
		for {
			try++
			if options.Retries <= try {
				err = errors.New("Exceeded maximum db connection retries")
				break
			}
			db, err = sql.Open("postgres", options.ConnectionURL)
			if err != nil {
				log.Printf("cannot connect to db, retrying...")
				select {
				case <-ctx.Done():
					break
				case <-time.After(options.Delay):
					continue
				}
			}
			break
		}
		connErrCh <- err
	}()
	select {
	case err = <-connErrCh:
		break
	case <-time.After(options.Timeout):
		return nil, errors.New("Failed connecting to database, timed out")
	case <-ctx.Done():
		return nil, errors.New("Failed connecting to database, context terminated")
	}
	return
}

func ConnectRabbitMQ(ctx context.Context, options ConnectionOptions) (conn *amqp.Connection, err error) {
	connErrCh := make(chan error, 1)
	defer close(connErrCh)
	go func() {
		try := 0
		for {
			try++
			if options.Retries <= try {
				err = errors.New("Exceeded maximum rabbit mq connection retries")
				break
			}
			conn, err = amqp.Dial(options.ConnectionURL)
			if err != nil {
				log.Printf("cannot connect to rabbit mq, retrying...")
				select {
				case <-ctx.Done():
					break
				case <-time.After(options.Delay):
					continue
				}
			}
			break
		}
		connErrCh <- err
	}()
	select {
	case err = <-connErrCh:
		break
	case <-time.After(options.Timeout):
		return nil, errors.New("Failed connecting to rabbit mq, timed out")
	case <-ctx.Done():
		return nil, errors.New("Failed connecting to rabbit mq, context terminated")
	}
	return
}
