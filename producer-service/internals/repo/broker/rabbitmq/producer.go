package rabbitmq

import (
	"github.com/kratos-14/pdf-compressor/producer-service/internals/model"
	amqp "github.com/rabbitmq/amqp091-go"
)

func (b *rabbitmqBroker) Produce(key string, msg []byte) error {
	rMQMsg := model.RabbitMQMessage{
		Exchange: "",
		RoutingKey: key,
		Mandatory: false,
		Immediate: false,
		Publishing: amqp.Publishing{
			ContentType: "application/json",
			Body: msg,
		},
	}
	return b.channel.Publish(rMQMsg.Exchange, rMQMsg.RoutingKey, rMQMsg.Mandatory, rMQMsg.Immediate, rMQMsg.Publishing)
}
