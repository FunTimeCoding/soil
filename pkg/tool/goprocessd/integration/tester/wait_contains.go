package tester

import (
	"strings"
	"testing"
	"time"
)

func (t *Tester) WaitContains(
	test *testing.T,
	substring string,
	command string,
	arguments ...string,
) {
	test.Helper()
	deadline := time.Now().Add(5 * time.Second)
	var last string

	for time.Now().Before(deadline) {
		last = t.Send(command, arguments...)

		if strings.Contains(last, substring) {
			return
		}

		time.Sleep(20 * time.Millisecond)
	}

	test.Fatalf(
		"timed out waiting for output containing %q, last output %q",
		substring,
		last,
	)
}
