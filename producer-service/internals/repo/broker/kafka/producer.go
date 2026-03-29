package kafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

func (b *kafkaBroker) Produce(topic string, msg []byte) error {
	mMsg := kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            []byte("msg"),
		Value:          msg,
	}
	return b.producer.Produce(&mMsg, nil)
}
