package base

import "testing"

func NewHidden(
	t *testing.T,
	tag string,
) *Server {
	t.Helper()

	return newWith(t, tag, "")
}
