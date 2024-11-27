package colt

import (
	"log"
)

type Logger interface {
	Print(v ...interface{})
	Panic(v ...interface{})
}

type DatabaseOption func(db *databaseOptions) error

// WithLogger sets the logger for the database. If none is provided, the default logger is used
func WithLogger(logger Logger) DatabaseOption {
	return func(db *databaseOptions) error {
		db.logger = logger
		return nil
	}
}

type databaseOptions struct {
	logger Logger
}

func (d *databaseOptions) Logger() Logger {
	if d.logger == nil {
		return log.Default()
	}
	return d.logger
}
