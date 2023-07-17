package controllers

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func getMongoClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	godotenv.Load(".env")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_ATLAS_URI")))
	checkFatality(err)
	checkFatality(client.Ping(ctx, readpref.Primary()))
	return client
}
