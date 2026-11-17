package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"testing"
)

func RefusingBravo(t *testing.T) *Tester {
	t.Helper()
	result := New(t, WithInstances("alfa", "bravo"))
	result.Bravo().Failure = validation.New("bad statement")

	return result
}
