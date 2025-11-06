package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

func (b *kafkaBroker) Produce(msg interface{}) error {
	mMsg := msg.(kafka.Message)
	return b.producer.Produce(&mMsg, nil)
}
