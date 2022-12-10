package colt

type BeforeInsertHook interface {
	BeforeInsert() error
}

type BeforeUpdateHook interface {
	BeforeUpdate() error
}