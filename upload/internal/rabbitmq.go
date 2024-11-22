package internal

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitClient struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func ConnectRabbitMQ(username, password, host, vhost string) (*amqp.Connection, error) {
	// Connection String format amqp://<username>:<password>@localhost:15672/<vhost>
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/%s", username, password, host, vhost))
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func CreateNewRabbitConn(conn *amqp.Connection) (RabbitClient, error) {
	ch, err := conn.Channel()
	if err != nil {
		return RabbitClient{}, err
	}

	return RabbitClient{
		conn: conn,
		ch:   ch,
	}, nil
}

func (rc RabbitClient) CreateNewRabbitQueue(queueName string, durable bool, autoDelete bool) (*amqp.Queue, error) {
	queue, err := rc.ch.QueueDeclare(queueName, durable, autoDelete, false, false, nil)
	if err != nil {
		return nil, err
	}

	return &queue, nil
}

func (rc RabbitClient) CloseRabbitConnection() error {
	return rc.ch.Close()
}
