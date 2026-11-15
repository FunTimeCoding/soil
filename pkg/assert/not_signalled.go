package assert

import (
	"testing"
	"time"
)

func NotSignalled[T any](
	t *testing.T,
	c <-chan T,
) {
	t.Helper()

	select {
	case <-c:
		t.Fatal("a signal came")
	case <-time.After(time.Second):
	}
}
