package utils

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB(ctx context.Context) (*mongo.Client, error) {
	mongoURL := os.Getenv("MONGO_URL")

	if mongoURL == "" {
		log.Fatal("MONGO_URL not set in .env file")
	}

	var client *mongo.Client
	var err error
	maxRetries := 10
	retryInterval := 15 * time.Second

	for i := 0; i < maxRetries; i++ {
        client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
        if err == nil {
            err = client.Ping(ctx, nil)
            if err == nil {
                log.Println("Connected to MongoDB!")
                return client, nil
            }
        }
        log.Printf("Failed to connect to MongoDB (attempt %d/%d): %v", i+1, maxRetries, err)
        time.Sleep(retryInterval)
    }

	log.Println("Connected to MongoDB!")
	return client, nil
}
