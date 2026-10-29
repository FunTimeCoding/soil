package tester

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
	"testing"
	"time"
)

func (t *Tester) WaitLog(
	test *testing.T,
	substring string,
	name string,
) {
	test.Helper()
	deadline := time.Now().Add(5 * time.Second)
	var last string

	for time.Now().Before(deadline) {
		last = join.NewLine(t.Log(name, false).Lines)

		if strings.Contains(last, substring) {
			return
		}

		time.Sleep(20 * time.Millisecond)
	}

	test.Fatalf(
		"timed out waiting for log containing %q, last %q",
		substring,
		last,
	)
}
