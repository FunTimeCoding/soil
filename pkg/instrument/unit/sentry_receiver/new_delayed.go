package sentry_receiver

import (
	"sync/atomic"
	"testing"
	"time"
)

func NewDelayed(
	t *testing.T,
	delay time.Duration,
) *atomic.Int32 {
	t.Helper()

	return build(t, delay)
}
