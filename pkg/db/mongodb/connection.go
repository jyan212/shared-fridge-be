package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongoDbClient struct {
	Client *mongo.Client
}

func NewConnection(uri string) (*MongoDbClient, error) {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("❌ Failed to connect to MongoDB")
		return nil, err
	}

	log.Println("✅ Connected to MongoDB")

	return &MongoDbClient{Client: client}, nil
}
