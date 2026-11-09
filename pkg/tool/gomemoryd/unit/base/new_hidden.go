package base

import (
	"testing"
	"time"
)

func NewHidden(
	t *testing.T,
	tag string,
) *Server {
	t.Helper()

	return serve(t, tag, "", time.Now)
}
