package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jensteichert/colt"
	"go.mongodb.org/mongo-driver/bson"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func initTracer() *sdktrace.TracerProvider {
	exporter, err := stdouttrace.New(
		stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatal(err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSyncer(exporter),
		sdktrace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("colt-example"),
				semconv.ServiceVersionKey.String("0.0.42"),
				semconv.DeploymentEnvironmentKey.String("local"),
			)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}))
	return tp
}

var tracer = initTracer().Tracer("colt-example")

type Database struct {
	Todos *colt.Collection[*Todo]
}

type Todo struct {
	colt.DocWithTimestamps `bson:",inline"`
	Title                  string `bson:"title" json:"title"`
}

func (t *Todo) BeforeInsert() error {
	t.DocWithTimestamps.BeforeInsert()
	fmt.Println("BeforeInsert executed")
	return nil
}

func initAndConnectToDbWithTracing() colt.Database {
	connectTraceCtx, span := tracer.Start(context.Background(), "coltInitDB")
	defer span.End()

	db := colt.NewDatabase().WithContext(connectTraceCtx)
	db.Connect("mongodb://localhost:27017/colt?readPreference=primary&directConnection=true&ssl=false", "colt")
	return db
}

func handleTodos(ctx context.Context) []*Todo {
	traceCtx, span := tracer.Start(ctx, "handleTodos")
	defer span.End()

	newTodo := Todo{Title: "Hello"}

	contextAwareTodoCollection := database.Todos.WithContext(traceCtx)

	todo, _ := contextAwareTodoCollection.Insert(&newTodo) // Will return a Todo
	insertedTodo, _ := contextAwareTodoCollection.FindById(todo.ID)

	fmt.Println(todo)

	contextAwareTodoCollection.UpdateById(todo.ID, todo)

	if insertedTodo != nil {
		fmt.Println(insertedTodo.Title)
	}

	allTodos, _ := contextAwareTodoCollection.Find(bson.M{"title": "Hello"})
	return allTodos
}

var database Database

func main() {
	db := initAndConnectToDbWithTracing()

	database = Database{
		Todos: colt.GetCollection[*Todo](&db, "todos"),
	}

	// This would usually be a request handler with its own request context
	allTodos := handleTodos(context.Background())

	for _, todo := range allTodos {
		fmt.Println(todo.ID)
	}
}
