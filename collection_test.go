package colt

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"math/rand"
	"testing"
	"time"
)

type testdoc struct {
	Doc   `bson:",inline"`
	Title string `bson:"title" json:"title"`
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
}
