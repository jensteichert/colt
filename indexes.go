package colt

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *Collection[T]) CreateIndex(keys bson.M) {
	go func() {
		mod := mongo.IndexModel{
			Keys: keys, Options: nil,
		}
		_, err := repo.collection.Indexes().CreateOne(DefaultContext(), mod)
		if err != nil {
		}
	}()
}
