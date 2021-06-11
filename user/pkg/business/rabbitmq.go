package business

import (
	"errors"

	amqp "github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}

func NewRabbitMQ(conn *amqp.Connection) (*RabbitMQ, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return &RabbitMQ{
		conn: conn,
		ch:   ch,
	}, nil
}

func (r *RabbitMQ) CreateQueue(name string) error {
	q, err := r.ch.QueueDeclare(
		name,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	r.q = q
	return nil
}

func (r *RabbitMQ) PublishMessage(message []byte) error {
	err := r.ch.Publish(
		"",
		r.q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
	if err != nil {
		return errors.New("Failed to publish message")
	}
	return nil
}
