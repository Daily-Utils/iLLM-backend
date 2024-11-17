package controllers

import (

	"go.mongodb.org/mongo-driver/mongo"
)

type Controller struct {
	MongoClient *mongo.Client
}
