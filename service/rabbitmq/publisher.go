package rabbitmq

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	Channel   *amqp.Channel
	QueueName string      `json:"queue_name"`
	Body      interface{} `json:"body"`
}

func (p *Publisher) Publish() error {
	q, err := p.Channel.QueueDeclare(
		p.QueueName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return err
	}
	body, _ := json.Marshal(p.Body)
	if err := p.Channel.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		}); err != nil {
		return err
	}
	return nil
}
