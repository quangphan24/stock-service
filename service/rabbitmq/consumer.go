package rabbitmq

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"log"
	"stock-service/payload"
	"stock-service/usecase"
	"stock-service/util"
)

var ListQueue = []string{
	util.UPDATE_QUANTITY_PRODUCT,
}

type Consumer struct {
	Conn    *amqp.Connection
	UseCase *usecase.UseCase
}

func NewConsumer(c *amqp.Connection, uc *usecase.UseCase) *Consumer {
	return &Consumer{
		Conn:    c,
		UseCase: uc,
	}
}
func (c *Consumer) StartConsumer() {
	for _, queue := range ListQueue {
		go c.Consume(queue)
	}
}

func (c *Consumer) Consume(queueName string) {
	ch, err := c.Conn.Channel()
	if err != nil {
		logrus.Error(err)
		return
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		logrus.Error(err)
		return
	}
	var forever chan struct{}

	go func() {
		for m := range msgs {
			log.Println(m)
			switch queueName {
			case util.UPDATE_QUANTITY_PRODUCT:
				var req payload.UpdateQuantity
				_ = json.Unmarshal(m.Body, &req)
				if err := req.Validate(*validator.New()); err != nil {
					logrus.Error(err)
					continue
				}
				err := c.UseCase.ProductUseCase.UpdateQuantity(req)
				if err != nil {
					logrus.Error(err)
				}
			}
		}
	}()
	<-forever
	return
}

func (c *Consumer) Close() {
	c.Conn.Close()
}
