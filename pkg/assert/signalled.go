package assert

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"testing"
	"time"
)

func Signalled[T any](
	t *testing.T,
	c <-chan T,
) T {
	t.Helper()

	select {
	case v := <-c:
		return v
	case <-time.After(soil.FixtureSignalWait):
		t.Fatal("the signal never came")
		var zero T

		return zero
	}
}
