package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Pictures of Athena!")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	athenaDatabase := client.Database("athena")
	picturesCollection := athenaDatabase.Collection("pictures")

	pic := bson.D{
		{"title", "Colors"},
		{"price", "19.99"},
		{"Artist", "Between The Buried and Me"},
	}
	insertOne(pic, picturesCollection, ctx)
}

func insertOne(pic bson.D, collection *mongo.Collection, ctx context.Context) int {
	picInsert, err := collection.InsertOne(ctx, pic)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(picInsert)
	return 0
}