package colt

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection[T Document] struct {
	collection *mongo.Collection
}

func (repo *Collection[T]) Insert(model T) (T, error) {
	if model.GetID() == "" {
		model.SetID(model.NewID())
	}

	if hook, ok := any(model).(BeforeInsertHook); ok {
		if err := hook.BeforeInsert(); err != nil {
			return model, err
		}
	}

	res, err := repo.collection.InsertOne(DefaultContext(), model)
	if err != nil && res != nil {
		model.SetID(res.InsertedID.(string))
	}

	return model, err
}

func (repo *Collection[T]) UpdateById(id string, model T) error {
	return repo.UpdateOne(bson.M{"_id": id}, model)
}

func (repo *Collection[T]) UpdateOne(filter interface{}, model T) error {
	if hook, ok := any(model).(BeforeUpdateHook); ok {
		if err := hook.BeforeUpdate(); err != nil {
			return err
		}
	}

	_, err := repo.collection.UpdateOne(DefaultContext(), filter, bson.M{"$set": model})
	return err
}

func (repo *Collection[T]) UpdateMany(filter interface{}, doc bson.M) error {
	_, err := repo.collection.UpdateMany(DefaultContext(), filter, doc)
	return err
}


func (repo *Collection[T]) DeleteById(id string) error {
	res, err := repo.collection.DeleteOne(DefaultContext(), bson.M{"_id": id})

	if err != nil {
		return err
	}

	if res.DeletedCount < 1 {
		return errors.New("could not delete")
	}

	return nil
}

func (repo *Collection[T]) FindById(id interface{}) (T, error) {
	return repo.FindOne(bson.M{"_id": id})
}

func (repo *Collection[T]) FindOne(filter interface{}) (T, error) {
	var target T
	err := repo.collection.FindOne(DefaultContext(), filter).Decode(&target)

	return target, err
}

func (repo *Collection[T]) Find(filter interface{}, opts ...*options.FindOptions) ([]T, error) {
	csr, err := repo.collection.Find(DefaultContext(), filter, opts...)
	if err != nil {
		return nil, err
	}
	var result = []T{}
	if err = csr.All(DefaultContext(), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *Collection[T]) CountDocuments(filter interface{}) (int64, error) {
	count, err := repo.collection.CountDocuments(DefaultContext(), filter)
	return count, err
}

func (repo *Collection[T]) NewId() primitive.ObjectID {
	return primitive.NewObjectID()
}
