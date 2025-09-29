package repo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func (r *repo) DownloadFile(id primitive.ObjectID) (string, *gridfs.DownloadStream, error) {
	downloadStream, err := r.bucket.OpenDownloadStream(id)
	if err != nil {
		return "", nil, err
	}
	defer downloadStream.Close()
	var fileName string
	fileInfo := bson.M{}
	err = r.bucket.GetFilesCollection().FindOne(context.Background(), bson.M{"_id": id}).Decode(&fileInfo)
	if err != nil {
		return "", nil, err
	}
	fileNameRaw, ok := fileInfo["filename"].(bson.Raw)
	if !ok {
		fileName = "compressed.pdf"
	} else {
		fileName = fileNameRaw.String()
	}
	return fileName, downloadStream, nil
}
