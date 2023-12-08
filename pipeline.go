package colt

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type Pipeline[T Document, R any] struct {
	collection *mongo.Collection
	pipeline   mongo.Pipeline
}

func (p *Pipeline[T, R]) Match(filter bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$match", filter}})
	return p
}

func (p *Pipeline[T, R]) Group(group bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$group", group}})
	return p
}

func (p *Pipeline[T, R]) Sort(sort bson.M) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, bson.D{{"$sort", sort}})
	return p
}

func (p *Pipeline[T, R]) Add(stage bson.D) *Pipeline[T, R] {
	p.pipeline = append(p.pipeline, stage)
	return p
}

func (p *Pipeline[T, R]) Run() (Cursor[R], error) {
	c, err := p.collection.Aggregate(DefaultContext(), p.pipeline)
	if err != nil {
		return nil, err
	}
	return &cursor[R]{&sync.Mutex{}, nil, nil, false, DefaultContext(), c}, nil
}
