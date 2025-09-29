package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func (s *service) DownloadFile(id string) (string, *gridfs.DownloadStream, error) {
	fileId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", nil, err
	}
	fileName, writer, err := s.repo.DownloadFile(fileId)
	if err != nil {
		return "", nil, err
	}
	return fileName, writer, nil
}
