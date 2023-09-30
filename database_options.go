package colt

import "log"

type Logger interface {
	Print(v ...interface{})
	Panic(v ...interface{})
}

type Option func(db *opts) error

func WithConnectionString(connectionString string) Option {
	return func(db *opts) error {
		db.connectionString = connectionString
		return nil
	}
}

func WithDBName(dbName string) Option {
	return func(db *opts) error {
		db.dbName = dbName
		return nil
	}
}

func WithLogger(logger Logger) Option {
	return func(db *opts) error {
		db.logger = logger
		return nil
	}
}

type opts struct {
	connectionString string
	dbName           string
	logger           Logger
}

func (d *opts) ConnectionString() string {
	return d.connectionString
}

func (d *opts) DBName() string {
	return d.dbName
}

func (d *opts) Logger() Logger {
	if d.logger == nil {
		return log.Default()
	}
	return d.logger
}
