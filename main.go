package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *UserCollection

func init() {
	// Connect to mongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	userCollection = &UserCollection{"clients", "users", nil}

	userCollection.Connect(client)
}

func main() {

	router := Router{"localhost", 3031}
	router.Int()
}
