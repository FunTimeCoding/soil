package base

import (
	"testing"
	"time"
)

func NewRooted(
	t *testing.T,
	root string,
) *Server {
	t.Helper()

	return serve(t, "", root, time.Now)
}
