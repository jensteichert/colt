package colt

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Document interface {
	SetID(id string)
	GetID() string

	NewID() string
	//CastID(id interface{}) (interface{}, error)
}

type Doc struct {
	ID string `bson:"_id,omitempty" json:"_id,omitempty"`
}

func (doc *Doc) NewID() string {
	return primitive.NewObjectID().Hex()
}

func (doc *Doc) SetID(id string) {
	doc.ID = id
}

func (doc *Doc) GetID() string {
	return doc.ID
}

type DocWithTimestamps struct {
	Doc       `bson:",inline"`
	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

func (doc *DocWithTimestamps) BeforeInsert() error {
	doc.CreatedAt = time.Now()
	return nil
}

func (doc *DocWithTimestamps) BeforeUpdate() error {
	now := time.Now()
	doc.UpdatedAt = &now
	return nil
}
