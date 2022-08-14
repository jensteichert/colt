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

type Database struct {
	Todos *colt.Collection[Todo]
}

func main() {
	db := colt.Database{}
	db.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")

	database := Database{
		Todos: colt.GetCollection[Todo](&db, "todos"),
	}

	todo := Todo{
		Title: "Hello",
	}

	database.Todos.Insert(todo)
	todos, _ := database.Todos.Find(bson.M{"title": "Hello"})

	for _, todo := range todos {
		fmt.Print(todo.Title)
	}
}
