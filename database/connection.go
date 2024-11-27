package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	Client   *mongo.Client
	Database *mongo.Database
	UserColl *mongo.Collection
)

func ConnectToMongoDB() {
	uri := os.Getenv("MONGO_URI")
	if uri == "" {
		log.Fatal("Mongo URI not set")
	}

	ctx := context.TODO()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	Database = client.Database("blood-bank-management")
	UserColl = Database.Collection("users")
	fmt.Println("Connected to MongoDB successfully")
}
