package colt

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCollection_CreateIndex(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")

	collection := &CollectionImpl[*testdoc]{collection: mockDb.db.Collection("testdocs")}

	var indxs = []interface{}{}
	indexCursor, _ := collection.collection.Indexes().List(DefaultContext())
	indexCursor.All(DefaultContext(), &indxs)

	indexCountBefore := len(indxs)

	err := collection.CreateIndex(bson.D{
		{fmt.Sprint(rand.Int()), 1},
	})

	assert.Nil(t, err)

	indexCursor2, _ := collection.collection.Indexes().List(DefaultContext())
	indexCursor2.All(DefaultContext(), &indxs)

	// new index
	assert.Equal(t, len(indxs), indexCountBefore+1)
}

func TestCollection_CreateMultiKeyIndex(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")

	collection := &CollectionImpl[*testdoc]{collection: mockDb.db.Collection("testdocs")}

	var indxs = []interface{}{}
	indexCursor, _ := collection.collection.Indexes().List(DefaultContext())
	indexCursor.All(DefaultContext(), &indxs)

	indexCountBefore := len(indxs)

	err := collection.CreateIndex(bson.D{
		{fmt.Sprint(rand.Int()), 1},
		{fmt.Sprint(rand.Int()), 1},
	})

	assert.Nil(t, err)

	indexCursor2, _ := collection.collection.Indexes().List(DefaultContext())
	indexCursor2.All(DefaultContext(), &indxs)

	// new index
	assert.Equal(t, len(indxs), indexCountBefore+1)
}
