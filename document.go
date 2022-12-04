package colt

type Document interface {
	SetID(id string)
	GetID() string
	//CastID(id interface{}) (interface{}, error)
}

type CDocument struct {
	ID    string `bson:"_id,omitempty" json:"_id,omitempty"`
}

func (f *CDocument) SetID(id string) {
	f.ID = id
}

func (f *CDocument) GetID() string {
	return f.ID
}

/*func (f *CDocument) CastID(id interface{}) (interface{}, error){
	if string, ok := id.(string); ok {
		return primitive.ObjectIDFromHex(string)
	}

	return id, nil
}*/