package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var CtxDB context.Context

var dBClient *mongo.Client
var DB *mongo.Database

var MoviesCollection *mongo.Collection

func InitDB() {
	log.Println("Establishing connection to Database @ " + os.Getenv("DB_URL"))
	CtxDB, _ = context.WithTimeout(context.TODO(), 20*time.Second)
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URL"))
	dbClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	dBClient = dbClient

	var result bson.M
	if err := dbClient.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}
	log.Println("Pinged your deployment. You are successfully connected to MongoDB!!!")
	DB = dbClient.Database("blitz")
	MoviesCollection = DB.Collection("movies")
}

func ShutDownDB() {
	dBClient.Disconnect(context.TODO())
}
