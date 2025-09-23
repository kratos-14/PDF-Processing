package service

import (
	"mime/multipart"

	"github.com/kratos-14/pdf-compressor/producer-service/internals/repo"
)

type Service interface {
	UploadFile(*multipart.Form) error
	DownloadFile()
}

type service struct {
	repo repo.Repo
}

func New(repo *repo.Repo) Service {
	return &service{
		repo: *repo,
	}
}
