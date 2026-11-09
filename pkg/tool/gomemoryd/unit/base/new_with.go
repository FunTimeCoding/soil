package base

import (
	"testing"
	"time"
)

func NewWith(
	t *testing.T,
	clock func() time.Time,
) *Server {
	t.Helper()

	return serve(t, "", "", clock)
}
