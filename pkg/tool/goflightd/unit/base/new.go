package base

import (
	"testing"
	"time"
)

func New(t *testing.T) *Server {
	t.Helper()

	return NewWith(t, time.Now)
}
