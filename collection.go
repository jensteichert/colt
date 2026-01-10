package colt

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection[T Document] struct {
	collection *mongo.Collection
	ctx        context.Context // optional context
}

func (repo Collection[T]) WithContext(ctx context.Context) *Collection[T] {
	repo.ctx = ctx
	return &repo
}

func (repo *Collection[T]) ctxOrDefault() (context.Context, context.CancelFunc) {
	if ctx := repo.ctx; ctx == nil {
		return defaultContext()
	}
	if _, hasDeadline := repo.ctx.Deadline(); !hasDeadline {
		ctxWithTimeout, cf := context.WithTimeout(repo.ctx, defaultTimeout)
		return ctxWithTimeout, cf
	}
	return repo.ctx, nil
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

	ctx, cancel := repo.ctxOrDefault()
	if cancel != nil {
		defer cancel()
	}
	res, err := repo.collection.InsertOne(ctx, model)
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

	ctx, cancel := repo.ctxOrDefault()
	if cancel != nil {
		defer cancel()
	}
	_, err := repo.collection.UpdateOne(ctx, filter, bson.M{"$set": model})
	return err
}

func (repo *Collection[T]) UpdateMany(filter interface{}, doc bson.M) error {
	ctx, cancel := repo.ctxOrDefault()
	if cancel != nil {
		defer cancel()
	}
	_, err := repo.collection.UpdateMany(ctx, filter, doc)
	return err
}

func (repo *Collection[T]) DeleteById(id string) error {
	ctx, cancel := repo.ctxOrDefault()
	if cancel != nil {
		defer cancel()
	}
	res, err := repo.collection.DeleteOne(ctx, bson.M{"_id": id})

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
	ctx, cancel := repo.ctxOrDefault()
	if cancel != nil {
		defer cancel()
	}
	var target T
	err := repo.collection.FindOne(ctx, filter).Decode(&target)

	return target, err
}

func (repo *Collection[T]) Find(filter interface{}, opts ...*options.FindOptions) ([]T, error) {
	ctx, cancel := repo.ctxOrDefault()
	if cancel != nil {
		defer cancel()
	}
	csr, err := repo.collection.Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	var result = []T{}
	if err = csr.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *Collection[T]) CountDocuments(filter interface{}) (int64, error) {
	ctx, cancel := repo.ctxOrDefault()
	if cancel != nil {
		defer cancel()
	}
	count, err := repo.collection.CountDocuments(ctx, filter)
	return count, err
}

func (repo *Collection[T]) Aggregate(pipeline mongo.Pipeline, opts ...*options.AggregateOptions) ([]bson.M, error) {
	ctx, cancel := repo.ctxOrDefault()
	if cancel != nil {
		defer cancel()
	}
	csr, err := repo.collection.Aggregate(ctx, pipeline, opts...)

	var result = []bson.M{}
	if err = csr.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *Collection[T]) Drop() error {
	ctx, cancel := repo.ctxOrDefault()
	if cancel != nil {
		defer cancel()
	}
	err := repo.collection.Drop(ctx)
	return err
}

func (repo *Collection[T]) NewId() primitive.ObjectID {
	return primitive.NewObjectID()
}
