package colt

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

var testCtx = context.WithValue(context.Background(), "test", "ctx")

func TestCollection_CreateIndex(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")

	collection := GetCollection[*testdoc](&mockDb, "testdocs")

	var indxs = []interface{}{}
	indexCursor, _ := collection.collection.Indexes().List(testCtx)
	indexCursor.All(testCtx, &indxs)

	indexCountBefore := len(indxs)

	err := collection.CreateIndex(bson.D{
		{fmt.Sprint(rand.Int()), 1},
	})

	assert.Nil(t, err)

	indexCursor2, _ := collection.collection.Indexes().List(testCtx)
	indexCursor2.All(testCtx, &indxs)

	// new index
	assert.Equal(t, len(indxs), indexCountBefore+1)
}

func TestCollection_CreateMultiKeyIndex(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")

	collection := GetCollection[*testdoc](&mockDb, "testdocs")

	var indxs = []interface{}{}
	indexCursor, _ := collection.collection.Indexes().List(testCtx)
	indexCursor.All(testCtx, &indxs)

	indexCountBefore := len(indxs)

	err := collection.CreateIndex(bson.D{
		{fmt.Sprint(rand.Int()), 1},
		{fmt.Sprint(rand.Int()), 1},
	})

	assert.Nil(t, err)

	indexCursor2, _ := collection.collection.Indexes().List(testCtx)
	indexCursor2.All(testCtx, &indxs)

	// new index
	assert.Equal(t, len(indxs), indexCountBefore+1)
}
