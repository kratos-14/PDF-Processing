package broker

import (
	"log"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var (
	kafkaProducer *kafka.Producer
	kafkaOnce     sync.Once
)

func KafkaConnect() (*kafka.Producer) {
	if kafkaProducer == nil {
		kafkaOnce.Do(func() {
			p, err := kafka.NewProducer(&kafka.ConfigMap{
				"bootstraps.servers": "kafka-release.default.svc.cluster.local:9092",
			})
			if err != nil {
				log.Fatal("failed to create kafka producer: ", err)
			}
			kafkaProducer = p
			log.Println("Kafka Producer Created Successfully")
		})
	}
	return kafkaProducer
}
