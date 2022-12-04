package colt

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Document interface {
	SetID(id interface{}) Document
}

type CDocument struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}

func (f CDocument) SetID(id interface{}) Document {
	f.ID = id.(primitive.ObjectID)
	return f
}