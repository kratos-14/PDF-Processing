package model

import amqp "github.com/rabbitmq/amqp091-go"

type RabbitMQMessage struct {
	Exchange   string
	RoutingKey string
	Mandatory  bool
	Immediate  bool
	Publishing amqp.Publishing
}
