package colt

import (
	"github.com/stretchr/testify/assert"
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