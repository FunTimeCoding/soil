package tester

import (
	"strings"
	"testing"
	"time"
)

func (t *Tester) WaitOutput(
	test *testing.T,
	expected string,
	command string,
	arguments ...string,
) {
	test.Helper()
	deadline := time.Now().Add(5 * time.Second)
	var last string

	for time.Now().Before(deadline) {
		last = strings.TrimSpace(t.Send(command, arguments...))

		if last == expected {
			return
		}

		time.Sleep(20 * time.Millisecond)
	}

	test.Fatalf("timed out waiting for %q, last output %q", expected, last)
}
