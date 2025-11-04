package broker

import (
	"log"
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

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

func CreateChannel() {
	conn := RabbitMQConnect()
	ch, err := conn.Channel()
	ch.Pu
}
