package wait_tester

import (
	"testing"
	"time"
)

func Until(
	t *testing.T,
	condition func() bool,
	wait time.Duration,
) {
	t.Helper()
	deadline := time.Now().Add(wait)

	for !condition() {
		if time.Now().After(deadline) {
			t.Fatal("condition never held")
		}

		time.Sleep(time.Millisecond)
	}
}
