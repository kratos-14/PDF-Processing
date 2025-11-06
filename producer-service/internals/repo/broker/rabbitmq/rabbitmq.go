package rabbitmq

import (
	"log"
	"sync"

	"github.com/kratos-14/pdf-compressor/producer-service/internals/repo/broker"
	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbitmqBroker struct {
	channel *amqp.Channel
}

func New(ch *amqp.Channel) broker.Broker {
	return &rabbitmqBroker{
		channel: ch,
	}
}

var (
	rabbitMQConnect *amqp.Connection
	rabbitMQOnce    sync.Once
)

func RabbitMQConnect() *amqp.Connection {
	if rabbitMQConnect == nil {
		rabbitMQOnce.Do(func() {
			rc, err := amqp.Dial("")
			if err != nil {
				log.Fatal("failed to connect to RabbitMQ: ", err)
			}
			rabbitMQConnect = rc
			log.Println("RabbitMq Connection Successfull")
		})
	}
	return rabbitMQConnect
}

func CreateChannel(conn *amqp.Connection) (*amqp.Channel) {
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed creating channel. error: %v", err)
	}
	return ch
}

func QueueDeclare(ch *amqp.Channel, queueName string) amqp.Queue {
	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("failed to declare queue. error: %v", err)
	}
	return q
}
