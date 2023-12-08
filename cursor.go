package colt

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type Cursor[T any] interface {
	// Next returns the next document in the cursor. If there is an error, it will be returned along with a nil document.
	// If there are no more documents, both the document and error will be nil, and the ok flag will be false.
	// Common usage:
	//
	//	for {
	//		doc, err, ok := cursor.Next()
	//		if !ok {
	//			break
	//		}
	//		if err != nil {
	//			// handle error
	//		}
	//		// do something with doc
	//	}
	Next() (*T, error, bool)

	// Value returns the current document in the cursor. If there was an error decoding the document, it will be returned.
	// If there are no more documents, both the document and error will be nil. It will not advance the cursor, so calling
	// Value() multiple times will return the same document, if Next() has been called initially. If Next() has not been
	// called, Value() will return nil, nil.
	Value() (*T, error)

	// Close closes the cursor. It should be called when the cursor is no longer needed. It is safe to call Close() multiple
	// times. It is not safe to call Next() or Value() after Close() has been called.
	Close() error
}

type cursor[T any] struct {
	lock       sync.Locker
	currentVal *T
	currentErr error
	isClosed   bool

	ctx    context.Context
	cursor *mongo.Cursor
}

func (p *cursor[T]) Next() (*T, error, bool) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.isClosed {
		return nil, nil, false
	}

	ok := p.cursor.Next(p.ctx)
	if !ok {
		p.currentVal = nil
		p.currentErr = nil
		return nil, nil, false
	}

	p.currentVal = new(T)
	p.currentErr = p.cursor.Decode(p.currentVal)
	return p.currentVal, p.currentErr, true
}

func (p *cursor[T]) Value() (*T, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if p.isClosed {
		return nil, nil
	}

	return p.currentVal, p.currentErr
}

func (p *cursor[T]) Close() error {
	p.lock.Lock()
	defer p.lock.Unlock()
	
	p.currentVal = nil
	p.currentErr = nil
	p.isClosed = true

	return p.cursor.Close(p.ctx)
}
