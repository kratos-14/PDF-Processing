package main

import (
	"log"
	"net/http"

	"github.com/kratos-14/pdf-compressor/producer-service/internals/handler"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/repo"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/repo/broker"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/server"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/service"
)

func main() {
	db, bucket := repo.MongoConnect()
	producer := broker.KafkaConnect()
	repo := repo.New(db, bucket)
	broker := broker.New(producer)
	service := service.New(&repo, &broker)
	handler := handler.New(service)
	mux := server.New(handler)
	newServer := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Fatal(newServer.ListenAndServe())
}
