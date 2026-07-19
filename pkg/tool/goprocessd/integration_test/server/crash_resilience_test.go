package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/integration_test/tester"
	"testing"
)

func TestProcessFailToStartDoesNotCrashManager(t *testing.T) {
	s := tester.New(
		t,
		"broken: /nonexistent/binary\nhealthy: sleep 60\n",
		"",
	)
	s.WaitOutput(t, "broken\n*healthy", "status")
}

func TestProcessCrashMidSessionDoesNotCrashManager(t *testing.T) {
	s := tester.New(
		t,
		"short: sleep 0.1\nhealthy: sleep 60\n",
		"",
	)
	s.WaitOutput(t, "short\n*healthy", "status")
}
