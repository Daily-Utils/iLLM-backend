package utils

import (
    "context"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() (*mongo.Client, error) {
    mongoURL := os.Getenv("MONGO_URL")

    log.Println("MONGO_URL: ", mongoURL)

    if mongoURL == "" {
        log.Fatal("MONGO_URL not set in .env file")
    }

    clientOptions := options.Client().ApplyURI(mongoURL)

    ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
    defer cancel() // Ensure the context is canceled to free resources

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        return nil, err
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        return nil, err
    }

    log.Println("Connected to MongoDB!")
    return client, nil
}