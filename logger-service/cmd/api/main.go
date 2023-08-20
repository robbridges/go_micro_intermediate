package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	webPORT  = "80"
	rpcPORT  = "5001"
	mongoURL = "mongodb://mongo:27017"
	gRPCPort = "50001"
)

var client *mongo.Client

type App struct {
}

func main() {
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panicln(err)
	}

	client = mongoClient

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func connectToMongo() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})
	c, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Println("Error connecting", err)
		return nil, err
	}
	return c, nil
}
