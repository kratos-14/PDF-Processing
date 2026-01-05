package main

import (
	"log"
	"net/http"

	"github.com/kratos-14/pdf-compressor/producer-service/internals/handler"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/repo"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/repo/broker/kafka"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/server"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/service"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/utils"
)

var (
	dbConnString       string
	kafkaConnString    string
	brokerType         string
	rabbitMQConnString string
	kafkaTopic         string
	rabbitMQQueue      string
)

func main() {
	db, bucket := repo.MongoConnect()
	producer := kafka.KafkaConnect()
	repo := repo.New(db, bucket)
	broker := kafka.New(producer)
	service := service.New(&repo, &broker)
	handler := handler.New(service)
	mux := server.New(handler)
	newServer := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Fatal(newServer.ListenAndServe())
}

func init() {
	getAllEnvVariables()
}

func getAllEnvVariables() {
	dbConnString = utils.GetEnv("DB_CONN_STR", "")
	if dbConnString == "" {
		log.Fatal("DB_CONN_STR is not set")
	}
	brokerType = utils.GetEnv("BROKER_TYPE", "")
	if brokerType == "" {
		log.Fatal("BROKER_TYPE is not set")
	}
	switch brokerType {
	case "kafka":
		kafkaConnString = utils.GetEnv("KAFKA_CONN_STR", "")
		if kafkaConnString == "" {
			log.Fatal("KAFKA_CONN_STR is not set")
		}
		kafkaTopic = utils.GetEnv("KAFKA_TOPIC", "")
		if kafkaTopic == "" {
			log.Fatal("KAFKA_TOPIC is not set")
		}
	case "rabbitmq":
		rabbitMQConnString = utils.GetEnv("RABBITMQ_CONN_STR", "")
		if rabbitMQConnString == "" {
			log.Fatal("RABBITMQ_CONN_STR is not set")
		}
		rabbitMQQueue = utils.GetEnv("RABBITMQ_QUEUE", "")
		if rabbitMQQueue == "" {
			log.Fatal("RABBITMQ_QUEUE is not set")
		}
	default:
		log.Fatal("BROKER_TYPE is not set or invalid")
	}

}
