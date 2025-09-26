package broker

import "github.com/confluentinc/confluent-kafka-go/kafka"

type Broker interface {
	Produce(kafka.Message) error
}

type broker struct {
	producer *kafka.Producer
}

func New(producer *kafka.Producer) Broker {
	return &broker{
		producer: producer,
	}
}
