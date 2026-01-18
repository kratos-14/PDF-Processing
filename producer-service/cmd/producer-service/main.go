package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kratos-14/pdf-compressor/producer-service/internals/handler"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/repo"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/repo/broker/rabbitmq"
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
	connection := rabbitmq.RabbitMQConnect()
	channel := rabbitmq.CreateChannel(connection)
	_ = rabbitmq.QueueDeclare(channel, rabbitMQQueue)
	repo := repo.New(db, bucket)
	broker := rabbitmq.New(channel)
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
	usernameBase64 := utils.GetEnv("MONGO_USERNAME", "")
	passwordBase64 := utils.GetEnv("MONGO_PASSWORD", "")
	username, err := utils.Base64Decode(usernameBase64)
	if err != nil {
		log.Fatalf("failed decoding base64 username. error: %v\n", err)
	}
	password, err := utils.Base64Decode(passwordBase64)
	if err != nil {
		log.Fatalf("failed decoding base64 password. error: %v\n", err)
	}
	connectionString := fmt.Sprintf("mongodb://%v:%v@mongodb.default.svc.cluster.local:27017/myFiles?authSource=admin&authMechanism=SCRAM-SHA-1", username, password)
	dbConnString = utils.GetEnv("DB_CONN_STR", connectionString)
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
