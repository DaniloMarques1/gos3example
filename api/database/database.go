package database

import (
	"context"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	once     sync.Once
	conError error
	client   *mongo.Client
)

func GetDbConnection() (*mongo.Client, error) {
	once.Do(func() {
		uri := os.Getenv("MONGO_URI")
		client, conError = mongo.Connect(
			context.Background(),
			options.Client().ApplyURI(uri),
		)
		if conError != nil {
			return
		}

		conError = client.Ping(context.Background(), nil)
		if conError != nil {
			return
		}
	})

	if conError != nil {
		return nil, conError
	}

	return client, nil
}
