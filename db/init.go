package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gopkg.in/mgo.v2/bson"
	"nc_student.com/v1/config"
)

var Client *mongo.Client

func Test() {
	fmt.Println("connect & insert db")
}

func init() {

	connect()
	insertNunber()
}

func insertNunber() {
	collection := Client.Database("testing").Collection("numbers")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, _ := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	id := res.InsertedID
	fmt.Println(id)
}

func connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Config.Mongo.Uri))
	if err != nil {
		log.Fatalf("connect error :%v", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())

	Client = client
}
