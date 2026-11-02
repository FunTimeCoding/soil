package unit

import (
	"testing"
	"time"
)

func await(
	t *testing.T,
	condition func() bool,
) {
	t.Helper()
	deadline := time.Now().Add(5 * time.Second)

	for !condition() {
		if time.Now().After(deadline) {
			t.Fatal("condition never held")
		}

		time.Sleep(time.Millisecond)
	}
}
