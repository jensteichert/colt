package colt

type BeforeInsertHook interface {
	BeforeInsert() error
}

type BeforeUpdate interface {
	BeforeUpdate() error
}