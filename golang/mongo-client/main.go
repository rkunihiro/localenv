package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	Isbn      string `json:"isbn"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

var mongoURI = os.Getenv("MONGODB_URI")

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})).With("module", "mongo-client")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Error("mongo.Connect failed", slog.Any("error", err))
		os.Exit(1)
	}
	log.Info("mongo.Connect success")

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Error("client.Disconnect failed", slog.Any("error", err))
			os.Exit(1)
		}
		log.Info("client.Disconnect success")
	}()

	if err := client.Ping(ctx, nil); err != nil {
		log.Error("client.Ping failed", slog.Any("error", err))
		os.Exit(1)
	} else {
		log.Info("client.Ping success")
	}

	database := client.Database("test")

	collection := database.Collection("books")

	book := &Book{}
	if err := collection.FindOne(context.TODO(), bson.D{{"isbn", "4873112699"}}).Decode(book); err != nil {
		log.Error("collection.FindOne failed", slog.Any("error", err))
		os.Exit(1)
	}
	log.Info("collection.FindOne success", slog.Any("book", book))
}
