package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"log-service/data"
	"net/http"

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
	Models data.Models
}

func main() {
	mongoClient, err := connectToMongo()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("Connected to Mongo")
	client = mongoClient

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := App{
		Models: data.New(client),
	}
	fmt.Println("Starting service on port", webPORT)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPORT),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Println("Error stating server: %s", err.Error())
	}

}

func (app *App) serve() {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPORT),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Println("Error stating server")
	}
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
