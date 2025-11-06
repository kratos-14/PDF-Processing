package rabbitmq

import "github.com/kratos-14/pdf-compressor/producer-service/internals/model"

func (b *rabbitmqBroker) Produce(msg interface{}) error {
	rMQMsg := msg.(model.RabbitMQMessage)
	return b.channel.Publish(rMQMsg.Exchange, rMQMsg.RoutingKey, rMQMsg.Mandatory, rMQMsg.Immediate, rMQMsg.Publishing)
}