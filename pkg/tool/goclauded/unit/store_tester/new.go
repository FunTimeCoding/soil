package store_tester

import (
	"testing"
	"time"
)

func New(t *testing.T) *Tester {
	t.Helper()

	return NewInZone(t, time.UTC)
}
