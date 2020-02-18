package mongo

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var instantiated *mongoDriver
var once sync.Once

func New() *mongoDriver {
	once.Do(func() {
		instantiated = &mongoDriver{}

		// Create client
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
		defer cancel()

		// Check if initial connection failed
		if err != nil {
			panic("Failed initial connection to DB")
		}

		// Connect to DB
		ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
		err = client.Ping(ctx, readpref.Primary())
		defer cancel()

		if err != nil {
			panic("Failed to connect to DB")
		}

	})
	return instantiated
}

type mongoDriver struct {
	O interface{}
}
