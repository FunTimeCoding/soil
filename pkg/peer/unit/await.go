package unit

import (
	"github.com/funtimecoding/soil/pkg/peer/constant"
	"testing"
	"time"
)

func await[T any](
	t *testing.T,
	c <-chan T,
) T {
	t.Helper()

	select {
	case v := <-c:
		return v
	case <-time.After(constant.FixtureSignalWait):
		t.Fatal("the signal never came")
		var zero T

		return zero
	}
}
