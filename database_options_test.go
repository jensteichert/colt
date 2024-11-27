package colt

import (
	"log"
	"reflect"
	"testing"
)

var (
	_ Logger = (*testLogger)(nil)
)

type testLogger struct{}

func (l *testLogger) Print(v ...interface{}) {}
func (l *testLogger) Panic(v ...interface{}) {}

func TestWithLogger(t *testing.T) {
	t.Run("should set the logger for the database", func(t *testing.T) {
		logger := &testLogger{}
		db := &Database{options: databaseOptions{}}
		db.applyOptions(WithLogger(logger))
		if !reflect.DeepEqual(db.options.logger, logger) {
			t.Errorf("WithLogger() = %v, want %v", db.options.logger, logger)
		}
	})
}

func Test_databaseOptions_Logger(t *testing.T) {
	t.Run("should return the default logger", func(t *testing.T) {
		want := log.Default()
		db := &Database{options: databaseOptions{}}

		if got := db.options.Logger(); !reflect.DeepEqual(want, got) {
			t.Errorf("WithLogger() = %v, want %v", got, want)
		}
	})

	t.Run("should return the provided logger from the database", func(t *testing.T) {
		want := &testLogger{}
		db := &Database{options: databaseOptions{
			logger: want,
		}}

		if got := db.options.Logger(); !reflect.DeepEqual(want, got) {
			t.Errorf("WithLogger() = %v, want %v", got, want)
		}
	})
}
