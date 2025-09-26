package broker

import "github.com/confluentinc/confluent-kafka-go/kafka"

func (b *broker) Produce(msg kafka.Message) error {
	return b.producer.Produce(&msg, nil)
}
