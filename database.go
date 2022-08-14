package colt

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Database struct {
	db     *mongo.Database
	client *mongo.Client
}

func (db *Database) ConnectWithDBName(dbName string) {
	dbURL := "mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false"
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbURL))
	if err != nil {
		log.Fatal(err)
	}
	db.client = client
	err = db.client.Ping(context.Background(), readpref.Primary())
	if err == nil {
		log.Print("Connected to MongoDB!")
	} else {
		log.Panic("Could not connect to MongoDB! Please check if mongo is running.", err)
	}
	db.db = db.client.Database(dbName)
}

func (db *Database) Connect() {
	db.ConnectWithDBName("colt")
}

func DefaultContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx
}

func GetRepo[T Document](db *Database, collectionName string) *Repo[T] {
	return &Repo[T]{db.db.Collection(collectionName)}
}
