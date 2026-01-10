package colt

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type testdoc struct {
	Doc   `bson:",inline"`
	Title string `bson:"title" json:"title"`
}

type testdocWithCustomID struct {
	Doc   `bson:",inline"`
	Title string `bson:"title" json:"title"`
}

func (pg *testdocWithCustomID) NewID() string {
	return "test_" + primitive.NewObjectID().Hex()
}

func TestCollection_Insert(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")

	collection := GetCollection[*testdoc](&mockDb, "testdocs")
	doc := testdoc{Title: fmt.Sprint(rand.Int())}

	inserted, err := collection.Insert(&doc)

	result, err := collection.FindById(inserted.ID)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, doc.ID, result.ID)

	mockDb.Disconnect()
}

func TestCollection_Insert_WithCustomId(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")

	collection := GetCollection[*testdocWithCustomID](&mockDb, "testdocs")
	doc := testdocWithCustomID{Title: fmt.Sprint(rand.Int())}

	inserted, err := collection.Insert(&doc)
	assert.True(t, strings.HasPrefix(doc.ID, "test_"))

	result, err := collection.FindById(inserted.ID)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, doc.ID, result.ID)
	assert.True(t, strings.HasPrefix(result.ID, "test_"))

	mockDb.Disconnect()
}

func TestCollection_Insert_WithError(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")

	collection := GetCollection[*testdocWithCustomID](&mockDb, "testdocs")
	doc := testdocWithCustomID{Title: fmt.Sprint(rand.Int())}

	inserted, err := collection.Insert(&doc)
	_, err2 := collection.Insert(&doc)

	assert.Nil(t, err)
	assert.NotNil(t, inserted)
	assert.NotNil(t, err2)

	mockDb.Disconnect()
}

func TestCollection_FindById(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")

	collection := GetCollection[*testdoc](&mockDb, "testdocs")
	doc := testdoc{Title: fmt.Sprint(rand.Int())}
	doc2 := testdoc{Title: "Test2"}

	collection.Insert(&doc)
	collection.Insert(&doc2)

	result, err := collection.FindById(doc.ID)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, doc.ID, result.ID)

	mockDb.Disconnect()
}

func TestCollection_FindOne(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")

	collection := GetCollection[*testdoc](&mockDb, "testdocs")
	doc := testdoc{Title: fmt.Sprint(rand.Int())}
	doc2 := testdoc{Title: "Test2"}

	collection.Insert(&doc)
	collection.Insert(&doc2)

	result, err := collection.FindOne(bson.M{"title": doc.Title})

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, doc.ID, result.ID)

	mockDb.Disconnect()
}

func TestCollection_Find(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")

	collection := GetCollection[*testdoc](&mockDb, "testdocs")

	title := fmt.Sprint(rand.Int())
	doc := testdoc{Title: title}
	doc2 := testdoc{Title: "Test2"}

	collection.Insert(&doc)
	collection.Insert(&doc2)

	result, err := collection.Find(bson.M{"title": title})

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(result), 1)
	assert.Equal(t, doc.ID, result[0].ID)

	mockDb.Disconnect()
}

func TestCollection_Find_Empty(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")

	collection := GetCollection[*testdoc](&mockDb, "testdocs")

	title := fmt.Sprint(rand.Int())
	doc := testdoc{Title: title}
	doc2 := testdoc{Title: "Test2"}

	collection.Insert(&doc)
	collection.Insert(&doc2)

	result, err := collection.Find(bson.M{"title": "NonExisting"})

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, len(result), 0)
	assert.Equal(t, result, []*testdoc{})

	mockDb.Disconnect()
}

func TestCollection_Find_WithError(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")
	collection := GetCollection[*testdoc](&mockDb, "testdocs")

	_, err := collection.Find(bson.M{"$exists": "NonExisting"})

	assert.NotNil(t, err)
}

func TestCollection_UpdateOne(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")

	collection := GetCollection[*testdoc](&mockDb, "testdocs")

	random := fmt.Sprint(rand.Int())
	doc := testdoc{Title: random}

	collection.Insert(&doc)

	fmt.Println(doc.ID)
	doc.Title = doc.Title + " updated"
	err := collection.UpdateOne(bson.M{"title": random}, &doc)

	totalWithNewTitle, err := collection.CountDocuments(bson.M{"title": random + " updated"})

	assert.Nil(t, err)
	assert.NotNil(t, totalWithNewTitle)
	assert.Equal(t, totalWithNewTitle, int64(1))

	mockDb.Disconnect()
}

func TestCollection_UpdateMany(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")

	collection := GetCollection[*testdoc](&mockDb, "testdocs")

	random := fmt.Sprint(rand.Int())
	doc := testdoc{Title: random}
	doc2 := testdoc{Title: random}

	collection.Insert(&doc)
	collection.Insert(&doc2)

	err := collection.UpdateMany(bson.M{"title": doc.Title}, bson.M{"$set": bson.M{"title": doc.Title + " updated"}})

	fmt.Println(err)
	time.Sleep(200 * time.Millisecond)

	totalWithNewTitle, err := collection.CountDocuments(bson.M{"title": doc.Title + " updated"})

	assert.Nil(t, err)
	assert.NotNil(t, totalWithNewTitle)
	assert.Equal(t, totalWithNewTitle, int64(2))

	mockDb.Disconnect()
}

// TODO
func TestCollection_DeleteById(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")

	collection := GetCollection[*testdoc](&mockDb, "testdocs")

	title := fmt.Sprint(rand.Int())
	doc := testdoc{Title: title}

	collection.Insert(&doc)
	err := collection.DeleteById(doc.ID)
	assert.Nil(t, err)

	result, err := collection.FindById(doc.ID)
	assert.Nil(t, result)
	assert.NotNil(t, err)

	collection.Drop()
	mockDb.Disconnect()
}

func TestCollection_CountDocuments(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")

	collection := GetCollection[*testdoc](&mockDb, "testdocs")

	title := fmt.Sprint(rand.Int())
	doc := testdoc{Title: title}
	doc2 := testdoc{Title: title}

	collection.Insert(&doc)
	collection.Insert(&doc2)

	result, err := collection.CountDocuments(bson.M{"title": title})

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result, int64(2))

	resultEmpty, err := collection.CountDocuments(bson.M{"title": "nonexistingtitle"})

	assert.Nil(t, err)
	assert.NotNil(t, resultEmpty)
	assert.Equal(t, resultEmpty, int64(0))

	mockDb.Disconnect()
}

func TestCollection_Aggregate(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")

	collection := GetCollection[*testdoc](&mockDb, "aggregatetest")

	title := fmt.Sprint(rand.Int())
	doc := testdoc{Title: title}
	doc2 := testdoc{Title: title}

	collection.Insert(&doc)
	collection.Insert(&doc2)

	result, err := collection.Aggregate(mongo.Pipeline{
		bson.D{
			{"$group", bson.D{
				{"_id", "$title"},
				{"count", bson.D{{"$sum", 1}}},
			}}}})

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result[0]["_id"], title)
	assert.Equal(t, result[0]["count"], int32(2))

	collection.Drop()
	mockDb.Disconnect()
}

func TestCollection_WithContext(t *testing.T) {
	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")

	collection := GetCollection[*testdoc](&mockDb, "testdocs")
	// Assert that the original collection has nil ctx (default)
	assert.Nil(t, collection.ctx)

	// Uses with explicit ctx
	ctxCollection := collection.WithContext(testCtx)

	// Assert that the underlying mongo collection pointer is identical
	assert.Equal(t, collection.collection, ctxCollection.collection)

	// Assert that the ctx is different
	assert.NotEqual(t, collection.ctx, ctxCollection.ctx)

	// Assert that the new collection has the context set
	assert.NotNil(t, ctxCollection.ctx)
	assert.Equal(t, testCtx, ctxCollection.ctx)

	secondCtx := context.WithValue(context.Background(), "test", "secondCtx")
	derivedCollection := ctxCollection.WithContext(secondCtx)

	// Assert that the underlying mongo collection pointer is identical
	assert.Equal(t, collection.collection, derivedCollection.collection)

	// Assert that the ctx is different
	assert.NotEqual(t, ctxCollection.ctx, derivedCollection.ctx)

	mockDb.Disconnect()
}

func TestCollection_ctxOrDefault(t *testing.T) {

	mockDb.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")

	// Assert that the original collection has nil ctx (default)
	staticCollection := GetCollection[*testdoc](&mockDb, "testdocs")
	defaultCtx, defaultCancelFunc := staticCollection.ctxOrDefault()
	assert.NotNil(t, defaultCtx)
	assert.NotNil(t, defaultCancelFunc)

	// Uses with explicit ctx
	ctxCollection := staticCollection.WithContext(testCtx)

	// Always has a deadline set
	defaultCtx, defaultCancelFunc = ctxCollection.ctxOrDefault()
	_, ok := defaultCtx.Deadline()
	assert.True(t, ok)
	// Enriches ctx with deadline and returns cancelFunc
	assert.NotNil(t, defaultCancelFunc)

	// Uses given ctx deadline if set
	deadline := time.Now().Add(42 * time.Second)
	deadlineCtx, cancel := context.WithDeadline(testCtx, deadline)

	dbWithManualCtxDeadline := staticCollection.WithContext(deadlineCtx)
	defaultCtx, defaultCancelFunc = dbWithManualCtxDeadline.ctxOrDefault()

	ctxDeadline, ok := deadlineCtx.Deadline()
	assert.True(t, ok)
	assert.Equal(t, ctxDeadline, deadline)
	// Doesnt return cancelFunc, as its managed by caller
	assert.Nil(t, defaultCancelFunc)

	cancel()
}
