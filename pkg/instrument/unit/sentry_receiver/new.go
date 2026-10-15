package sentry_receiver

import (
	"sync/atomic"
	"testing"
)

func New(t *testing.T) *atomic.Int32 {
	t.Helper()

	return build(t, 0)
}
