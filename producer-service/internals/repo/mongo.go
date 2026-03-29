package repo

import (
	"context"
	"fmt"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	mongoOnce sync.Once
	mongoDB   *mongo.Database
	bucket    *gridfs.Bucket
)

func MongoConnect(username, password string) (*mongo.Database, *gridfs.Bucket) {
	if mongoDB == nil {
		mongoOnce.Do(
			func() {
				// usernameBase64 := os.Getenv("MONGO_USERNAME")
				// passwordBase64 := os.Getenv("MONGO_PASSWORD")
				// usernameByte, _ := base64.StdEncoding.DecodeString(usernameBase64)
				// passwordByte, _ := base64.StdEncoding.DecodeString(passwordBase64)
				// username := string(usernameByte)
				// password := string(passwordByte)
				connectionString := fmt.Sprintf("mongodb://%v:%v@mongodb.default.svc.cluster.local:27017/myFiles?authSource=admin&authMechanism=SCRAM-SHA-1", username, password)
				clientOptions := options.Client().ApplyURI(connectionString)

				ctx := context.Background()
				mongoConnect, err := mongo.Connect(ctx, clientOptions)
				if err != nil {
					log.Fatal(err, " can't connect to mongodb altas")
				}
				err = mongoConnect.Ping(ctx, readpref.Primary())
				if err != nil {
					log.Fatal(err, " error pinging mongodb Atlas")
				}
				mongoConnection := mongoConnect
				mongoDB = mongoConnection.Database("myFiles")
				bucket, err = gridfs.NewBucket(mongoDB)
				if err != nil {
					log.Fatal(err, " error creating gridfs bucket")
				}
				log.Print("MongoDB Connected")
			})
	}
	return mongoDB, bucket
}
