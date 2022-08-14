# Colt
The mongodb ODM for Go i've always wanted

![Build](https://github.com/jensteichert/webvitals_exporter/workflows/Build/badge.svg)
![CodeQL](https://github.com/jensteichert/webvitals_exporter/workflows/CodeQL/badge.svg)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/jensteichert/colt)
<a href="https://github.com/jensteichert/colt/releases"><img src="https://img.shields.io/github/v/release/jensteichert/colt" /></a>


Colt leverages Generics to provide type-safe methods and decoding of documents. It therefor requires [Go 1.18+](https://tip.golang.org/doc/go1.18). 
### Installation

```
go get github.com/jensteichert/colt
```

### Quick Start
```golang
package main

import (
	"fmt"
	"github.com/jensteichert/colt"
	"go.mongodb.org/mongo-driver/bson"
)

type Todo struct {
	Id    string `bson:"_id,omitempty" json:"_id,omitempty"`
	Title string `bson:"title" json:"title"`
}

func main() {
	db := colt.Database{}
	db.Connect("mongodb://...", "myDatabaseName")

	todosCollection = colt.GetCollection[Todo](&db, "todos"),
	
	newTodo := Todo{
		Title: "Hello",
	}

	todosCollection.Insert(newTodo)

	todos, err := todosCollection.Find(bson.M{"title": "Hello"})

	for _, todo := range todos {
		fmt.Print(todo.Title)
	}
}
```