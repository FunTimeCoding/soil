//go:build local

package server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/integration_test/tester"
	"strings"
	"testing"
	"time"
)

func TestProcessFailToStartDoesNotCrashManager(t *testing.T) {
	s := tester.New(
		t,
		"broken: /nonexistent/binary\nhealthy: sleep 60\n",
		"",
	)
	time.Sleep(200 * time.Millisecond)
	result := s.Send("status")
	lines := strings.Split(result, "\n")
	assert.Integer(t, 2, len(lines))
	assert.String(t, " broken", lines[0])
	assert.String(t, "*healthy", lines[1])
}

func TestProcessCrashMidSessionDoesNotCrashManager(t *testing.T) {
	s := tester.New(
		t,
		"short: sleep 0.1\nhealthy: sleep 60\n",
		"",
	)
	time.Sleep(500 * time.Millisecond)
	result := s.Send("status")
	lines := strings.Split(result, "\n")
	assert.Integer(t, 2, len(lines))
	assert.String(t, " short", lines[0])
	assert.String(t, "*healthy", lines[1])
}
