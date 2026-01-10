package colt

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repo *Collection[T]) CreateIndex(keys bson.D) error {
	ctx, cancel := repo.ctxOrDefault()
	if cancel != nil {
		defer cancel()
	}
	mod := mongo.IndexModel{
		Keys:    keys,
		Options: nil,
	}
	_, err := repo.collection.Indexes().CreateOne(ctx, mod)

	return err
}
