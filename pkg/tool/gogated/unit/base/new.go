package base

import "testing"

func New(t *testing.T) *Server {
	t.Helper()

	return NewWith(t, nil)
}
