# Colt
The mongodb ODM for Go i've always wanted

![Build & Tests](https://github.com/jensteichert/webvitals_exporter/workflows/Build/badge.svg)
![CodeQL](https://github.com/jensteichert/colt/workflows/CodeQL/badge.svg)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/jensteichert/colt)
<a href="https://github.com/jensteichert/colt/releases"><img src="https://img.shields.io/github/v/release/jensteichert/colt" /></a>
[![Go Report Card](https://goreportcard.com/badge/github.com/jensteichert/colt)](https://goreportcard.com/report/github.com/jensteichert/colt)
[![Coverage Status](https://coveralls.io/repos/github/jensteichert/colt/badge.svg?branch=main)](https://coveralls.io/github/jensteichert/colt?branch=main)

Colt leverages Generics to provide type-safe methods and decoding of documents. It therefor requires [Go 1.18+](https://tip.golang.org/doc/go1.18). 
### Installation
To install Colt, use `go get`:
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

type Database struct {
	Todos *colt.Collection[*Todo]
}

type Todo struct {
	colt.DocWithTimestamps `bson:",inline"`
	Title string `bson:"title" json:"title"`
}

func main() {
	db := colt.Database{}
	db.Connect("mongodb://...", "myDatabaseName")

	database := Database{
		Todos: colt.GetCollection[*Todo](&db, "todos"),
	}

	newTodo := Todo{Title: "Hello"}

	todo, _ := database.Todos.Insert(&newTodo) // Will return a Todo
	insertedTodo, _ := database.Todos.FindById(todo.ID)

	allTodos, _ := database.Todos.Find(bson.M{"title": "Hello"})
}
```

## Features

### Hooks

#### ``BeforeInsert`` Hook
Triggers before a document will be inserted
```golang
type Todo struct {
	colt.DocWithTimestamps `bson:",inline"`
}

func(t *Todo) BeforeInsert() error {
	t.DocWithTimestamps.BeforeInsert()

        // Do something with t here
	return nil
}
```

#### ``BeforeUpdate`` Hook
Triggers before a document will be updated
```golang
func(t *Todo) BeforeUpdate() error {
	t.DocWithTimestamps.BeforeUpdate()

        // Do something with t here
	return nil
}
```


### ToDo
- [x] CRUD
- [x] Hooks
- [ ] Pagination
- [ ] Aggregations
- [ ] Transactions


