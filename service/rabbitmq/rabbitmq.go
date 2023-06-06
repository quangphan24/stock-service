package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func NewRabbitmq() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return nil, err
	}
	return conn.Channel()
}
