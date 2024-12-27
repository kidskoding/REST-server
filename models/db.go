package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client
var db *mongo.Database

func ConnectToDB() {
	var err error
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database("pizzastore")

	fmt.Println("Connected to MongoDB!")
}

func GetCollection() *mongo.Collection {
	if db == nil {
		log.Fatal("Database is not connected.")
	}
	return db.Collection("orders")
}
