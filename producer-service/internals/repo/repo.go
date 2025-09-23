package repo

import (
	"mime/multipart"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type Repo interface {
	UploadFile([]*multipart.FileHeader) error
}

type repo struct {
	db     *mongo.Database
	bucket *gridfs.Bucket
}

func New(db *mongo.Database, bucket *gridfs.Bucket) Repo {
	return &repo{
		db:     db,
		bucket: bucket,
	}
}
