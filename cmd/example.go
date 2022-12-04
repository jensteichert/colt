package main

import (
	"fmt"
	"github.com/jensteichert/colt"
	"go.mongodb.org/mongo-driver/bson"
)

type Database struct {
	Todos *colt.Collection[Todo]
}
type Todo struct {
	colt.CDocument `bson:",inline"`
	Title string `bson:"title" json:"title"`
}

func main() {
	db := colt.Database{}
	db.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")

	database := Database{
		Todos: colt.GetCollection[Todo](&db, "todos"),
	}

	new := Todo{
		Title: "Hello",
	}
	todo, _ := database.Todos.Insert(new)
	fmt.Println(todo.Title, todo.ID)

	// Should fails
	//todo, _ := database.Todos.Insert(newUser)


	todos, _ := database.Todos.Find(bson.M{"title": "Hello"})

	for _, todo := range todos {
		fmt.Println(todo.ID)
	}
}
