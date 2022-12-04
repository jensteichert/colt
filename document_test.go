package colt

import "testing"

type Todo struct {
	CDocument `bson:",inline"`
	Title string `bson:"title" json:"title"`
}
func TestCDocument_SetID(t *testing.T) {
	doc := Todo{}
	doc.SetID("638cda03871d719a9020c855")

	if doc.ID != "638cda03871d719a9020c855" {
		t.Errorf("ID is %s, should be 638cda03871d719a9020c855", doc.ID)
	}
}

func TestCDocument_GetID(t *testing.T) {
	doc := Todo{}

	if doc.GetID() != "" {
		t.Errorf("ID should be empty but is %s", doc.ID)
	}

	doc.SetID("638cda03871d719a9020c855")

	if doc.GetID() != "638cda03871d719a9020c855" {
		t.Errorf("ID should be 638cda03871d719a9020c855 but is %s", doc.ID)
	}
}