package server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/integration_test/tester"
	"strings"
	"testing"
)

func TestLogCapturesOutput(t *testing.T) {
	s := tester.New(
		t,
		"alfa: sh -c 'echo hello-from-alfa && sleep 60'\n",
		"",
	)
	s.WaitContains(t, "hello-from-alfa", "log", "alfa")
}

func TestLogUnknownProcess(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	result := s.Send("log", "nonexistent")
	assert.True(t, strings.HasPrefix(result, "error:"))
}
