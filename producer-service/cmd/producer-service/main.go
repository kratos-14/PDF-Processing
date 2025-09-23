package main

import (
	"log"
	"net/http"

	"github.com/kratos-14/pdf-compressor/producer-service/internals/handler"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/repo"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/server"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/service"
)

func main() {
	db, bucket, err := repo.MongoConnect()
	if err != nil {
		log.Fatal("Database Connection Failed")
	}
	repo := repo.New(db, bucket)
	service := service.New(&repo)
	handler := handler.New(service)
	mux := server.New(handler)
	newServer := http.Server{
		Addr: ":8080",
		Handler: mux,
	}
	log.Fatal(newServer.ListenAndServe())
}