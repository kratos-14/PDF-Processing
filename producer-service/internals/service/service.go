package service

import (
	"mime/multipart"

	"github.com/kratos-14/pdf-compressor/producer-service/internals/repo"
	"github.com/kratos-14/pdf-compressor/producer-service/internals/repo/broker"
)

type Service interface {
	UploadFile(*multipart.Form) error
	DownloadFile()
}

type service struct {
	repo repo.Repo
	broker broker.Broker
}

func New(repo *repo.Repo, broker *broker.Broker) Service {
	return &service{
		repo: *repo,
		broker: *broker,
	}
}
