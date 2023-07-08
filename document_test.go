package colt

import (
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strings"
	"testing"
)

type todo struct {
	Doc   `bson:",inline"`
	Title string `bson:"title" json:"title"`
}

func TestCDocument_SetID(t *testing.T) {
	doc := todo{}
	doc.SetID("638cda03871d719a9020c855")

	assert.Equal(t, doc.ID, "638cda03871d719a9020c855")
}

func TestCDocument_GetID(t *testing.T) {
	doc := todo{}
	assert.Empty(t, doc.GetID())

	doc.SetID("638cda03871d719a9020c855")

	assert.Equal(t, doc.ID, "638cda03871d719a9020c855")
}

func TestCDocument_NewID(t *testing.T) {
	doc := todoWithCustomId{}
	assert.Empty(t, doc.GetID())

	assert.NotEmpty(t, doc.NewID())
}

type todoWithCustomId struct {
	Doc   `bson:",inline"`
	Title string `bson:"title" json:"title"`
}

func (t *todoWithCustomId) NewID() string {
	return "td_" + primitive.NewObjectID().Hex()
}

func TestCDocument_NewID_Custom(t *testing.T) {
	doc := todoWithCustomId{}
	assert.Empty(t, doc.GetID())

	assert.NotEmpty(t, doc.NewID())
	assert.True(t, strings.HasPrefix(doc.NewID(), "td_"))
}