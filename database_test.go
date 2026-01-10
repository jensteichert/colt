package colt

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDatabase_WithContext(t *testing.T) {
	staticDb := NewDatabase()

	assert.Nil(t, staticDb.ctx)

	ctxAwareDb := staticDb.WithContext(testCtx)

	assert.NotEqual(t, staticDb, ctxAwareDb)
	assert.Equal(t, staticDb.db, ctxAwareDb.db)
	assert.Equal(t, staticDb.client, ctxAwareDb.client)
	assert.NotEqual(t, staticDb.ctx, ctxAwareDb.ctx)
}

func TestDatabase_ctxOrDefault(t *testing.T) {

	// Static, non explicit ctx use
	staticDb := NewDatabase()
	defaultCtx, defaultCancelFunc := staticDb.ctxOrDefault()
	assert.NotNil(t, defaultCtx)
	assert.NotNil(t, defaultCancelFunc)

	// Manuall, explicit ctx uses
	ctxDb := NewDatabase().WithContext(testCtx)

	// Always has a deadline set
	defaultCtx, defaultCancelFunc = ctxDb.ctxOrDefault()
	_, ok := defaultCtx.Deadline()
	assert.True(t, ok)
	// Enriches ctx with deadline and returns cancelFunc
	assert.NotNil(t, defaultCancelFunc)

	// Uses given ctx deadline if set
	deadline := time.Now().Add(42 * time.Second)
	deadlineCtx, cancel := context.WithDeadline(testCtx, deadline)
	dbWithManualCtxDeadline := NewDatabase().WithContext(deadlineCtx)
	defaultCtx, defaultCancelFunc = dbWithManualCtxDeadline.ctxOrDefault()

	ctxDeadline, ok := deadlineCtx.Deadline()
	assert.True(t, ok)
	assert.Equal(t, ctxDeadline, deadline)
	// Doesnt return cancelFunc, as its managed by caller
	assert.Nil(t, defaultCancelFunc)

	cancel()
}
