package repo

import (
	"io"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) UploadFile(files []*multipart.FileHeader) error {
	fileIds := make(map[primitive.ObjectID]string)
	for _, file := range files {
		fileHandle, err := file.Open()
		if err != nil {
			return err
		}
		defer fileHandle.Close()
		opts := options.GridFSUpload().SetMetadata(map[string]string{"originalFileName": file.Filename})
		fileID := primitive.NewObjectID()
		fileIds[fileID] = file.Filename
		uploadStream, err := r.bucket.OpenUploadStreamWithID(fileID, file.Filename, opts)
		if err != nil {
			return err
		}
		defer uploadStream.Close()
		_, err = io.Copy(uploadStream, fileHandle)
		if err != nil {
			return err
		}
	}
	return nil
}
