package assert

import (
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"testing"
)

func SkipOnOutage(
	t testing.TB,
	e error,
) {
	if connection.Classify(e) != nil {
		t.Helper()
		t.Skip(e)
	}
}
