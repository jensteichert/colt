package colt

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type testDoc struct {
	mock.Mock         `bson:"-"`
	DocWithTimestamps `bson:",inline"`
	Title             string `bson:"title" json:"title"`
}

func (t *testDoc) BeforeInsert() error {
	t.DocWithTimestamps.BeforeInsert()
	args := t.Called()
	return args.Error(0)
}

func (t *testDoc) BeforeUpdate() error {
	t.DocWithTimestamps.BeforeUpdate()
	args := t.Called()
	return args.Error(0)
}

var mockDb = Database{}

func TestBeforeInsertHook(t *testing.T) {
	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")
	doc := testDoc{Title: "Test"}

	doc.On("BeforeInsert").Return(nil)

	_, insertErr := GetCollection[*testDoc](&mockDb, "testdocs").Insert(&doc)
	assert.Nil(t, insertErr)

	doc.AssertExpectations(t)
}

func TestBeforeUpdateHook(t *testing.T) {
	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")
	doc := testDoc{Title: "Test"}

	doc.On("BeforeUpdate").Return(nil)

	updateErr := GetCollection[*testDoc](&mockDb, "testdocs").UpdateById(doc.ID, &doc)
	assert.Nil(t, updateErr)

	doc.AssertExpectations(t)
}
