package colt

import (
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
