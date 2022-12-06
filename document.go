package colt

import "time"

type Document interface {
	SetID(id string)
	GetID() string
	//CastID(id interface{}) (interface{}, error)
}

type Doc struct {
	ID    string `bson:"_id,omitempty" json:"_id,omitempty"`
}

func (doc *Doc) SetID(id string) {
	doc.ID = id
}

func (doc *Doc) GetID() string {
	return doc.ID
}

type DocWithTimestamps struct {
	Doc `bson:",inline"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

func (doc *DocWithTimestamps) BeforeInsert() error {
	doc.CreatedAt = time.Now()
	return nil
}

func (doc *DocWithTimestamps) BeforeUpdate() error {
	doc.UpdatedAt = time.Now()
	return nil
}