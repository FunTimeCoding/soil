package tester

import (
	"testing"
	"time"
)

func (t *Tester) Wait(
	test *testing.T,
	expected string,
	read func() string,
) {
	test.Helper()
	deadline := time.Now().Add(5 * time.Second)
	var last string

	for time.Now().Before(deadline) {
		last = read()

		if last == expected {
			return
		}

		time.Sleep(20 * time.Millisecond)
	}

	test.Fatalf("timed out waiting for %q, last %q", expected, last)
}
