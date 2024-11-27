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
		dbOptions := &databaseOptions{}
		WithLogger(logger)(dbOptions)
		if !reflect.DeepEqual(dbOptions.logger, logger) {
			t.Errorf("WithLogger() = %v, want %v", dbOptions.logger, logger)
		}
	})
}
