package base

import "testing"

func NewRooted(
	t *testing.T,
	root string,
) *Server {
	t.Helper()

	return newWith(t, "", root)
}
