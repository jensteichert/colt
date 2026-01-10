package colt

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

var defaultTimeout = 5 * time.Second

type Database struct {
	db     *mongo.Database
	client *mongo.Client
	ctx    context.Context // optional context
}

func NewDatabase() *Database {
	return &Database{}
}

func (db Database) WithContext(ctx context.Context) Database {
	db.ctx = ctx
	return db
}

func (db *Database) connect(options *options.ClientOptions, dbName string) error {
	ctx, cancel := db.ctxOrDefault()
	if cancel != nil {
		defer cancel()
	}

	client, err := mongo.Connect(ctx, options)
	if err != nil {
		log.Fatal(err)
	}
	db.client = client
	err = db.client.Ping(ctx, readpref.Primary())
	if err == nil {
		log.Print("Connected to MongoDB!")
	} else {
		log.Panic("Could not connect to MongoDB! Please check if mongo is running.", err)
		return err
	}
	db.db = db.client.Database(dbName)
	return nil
}

func (db *Database) Connect(connectionString string, dbName string) error {
	options := options.Client().ApplyURI(connectionString)
	options.Monitor = otelmongo.NewMonitor()
	err := db.connect(options, dbName)
	return err
}

func (db *Database) ctxOrDefault() (context.Context, context.CancelFunc) {
	if ctx := db.ctx; ctx == nil {
		return defaultContext()
	}
	if _, hasDeadline := db.ctx.Deadline(); !hasDeadline {
		ctxWithTimeout, cf := context.WithTimeout(db.ctx, defaultTimeout)
		return ctxWithTimeout, cf
	}
	return db.ctx, nil
}

func (db *Database) Ping() error {
	ctx, cancel := db.ctxOrDefault()
	if cancel != nil {
		defer cancel()
	}
	return db.client.Ping(ctx, nil)
}

func (db *Database) Disconnect() error {
	ctx, cancel := db.ctxOrDefault()
	if cancel != nil {
		defer cancel()
	}
	err := db.client.Disconnect(ctx)
	db.db = nil
	return err
}

func defaultContext() (context.Context, context.CancelFunc) {
	ctx, cf := context.WithTimeout(context.Background(), defaultTimeout)
	return ctx, cf
}

func GetCollection[T Document](db *Database, collectionName string) *Collection[T] {
	return &Collection[T]{db.db.Collection(collectionName), nil}
}
