package base

import (
	"testing"
	"time"
)

func New(t *testing.T) *Server {
	t.Helper()

	return serve(t, "", "", time.Now)
}
