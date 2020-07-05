package src

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var Client *mongo.Client

func ConnectDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://ip.alis.cetacis.dev:3000").SetMaxPoolSize(20))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = Client.Ping(ctx, readpref.Primary())
	fmt.Println("Connect open")
}

func DisconnectDb() {
	if err := Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	fmt.Println("Connect closed")
}