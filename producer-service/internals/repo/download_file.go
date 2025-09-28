package repo

import "go.mongodb.org/mongo-driver/bson/primitive"

func DownloadFile(id string) {
	fileId, err := primitive.ObjectIDFromHex(id)
}
