package main

import (
	"encoding/base64"
	"log"
	"net/http"

	"github.com/kratos-14/pdf-compressor/producer-service/internals/handler"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/repo"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/repo/broker"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/repo/broker/kafka"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/repo/broker/rabbitmq"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/server"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/service"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/utils"
)

var (
	username string
	password string
	// dbConnString       string
	kafkaConnString    string
	brokerType         string
	rabbitMQConnString string
	kafkaTopic         string
	rabbitMQQueue      string
)

func main() {
	db, bucket := repo.MongoConnect(username, password)
	var broker broker.Broker
	var topic string
	switch brokerType {
	case "rabbitmq":
		topic = rabbitMQQueue
		conn := rabbitmq.RabbitMQConnect(rabbitMQConnString)
		producer := rabbitmq.CreateChannel(conn)
		broker = rabbitmq.New(producer)
	case "kafka":
		topic = kafkaTopic
		producer := kafka.KafkaConnect(kafkaConnString)
		broker = kafka.New(producer)
	default:
		log.Fatalf("Invalid broker type %s specified\n", brokerType)
	}
	repo := repo.New(db, bucket)
	service := service.New(topic, &repo, &broker)
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
	// usernameBase64 := os.Getenv("MONGO_USERNAME")
	// passwordBase64 := os.Getenv("MONGO_PASSWORD")
	usernameBase64 := utils.GetEnv("MONGO_USERNAME", "")
	if usernameBase64 == "" {
		log.Fatal("MONGO_USERNAME is not set")
	}
	usernameByte, err := base64.StdEncoding.DecodeString(usernameBase64)
	if err != nil {
		log.Fatalf("failed decoding MONGO_USERNAME. error: %v\n", err)
	}
	passwordBase64 := utils.GetEnv("MONGO_PASSWORD", "")
	if passwordBase64 == "" {
		log.Fatal("MONGO_PASSWORD is not set")
	}
	passwordByte, err := base64.StdEncoding.DecodeString(passwordBase64)
	if err != nil {
		log.Fatalf("failed decoding MONGO_PASSWORD. error: %v\n", err)
	}
	username = string(usernameByte)
	password = string(passwordByte)
	// dbConnString = utils.GetEnv("DB_CONN_STR", "")
	// if dbConnString == "" {
	// 	log.Fatal("DB_CONN_STR is not set")
	// }
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
		rabbitMQConnString = utils.GetEnv("RABBITMQ_CONN_STR", "amqp://guest:guest@rabbitmq-service.default.svc.cluster.local:5672/")
		if rabbitMQConnString == "" {
			log.Fatal("RABBITMQ_CONN_STR is not set")
		}
		rabbitMQQueue = utils.GetEnv("RABBITMQ_QUEUE", "my-topic")
		if rabbitMQQueue == "" {
			log.Fatal("RABBITMQ_QUEUE is not set")
		}
	default:
		log.Fatal("BROKER_TYPE is not set or invalid")
	}

}
