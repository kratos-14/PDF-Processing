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

}
