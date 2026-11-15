package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"testing"
)

func RefusingBravo(t *testing.T) *Tester {
	t.Helper()
	s, clients := New(t, "alfa", "bravo")
	clients["bravo"].Failure = validation.New("bad statement")

	return s
}
