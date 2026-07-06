//go:build local

package server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/integration_test/tester"
	"strings"
	"testing"
	"time"
)

func TestLogCapturesOutput(t *testing.T) {
	s := tester.New(
		t,
		"alfa: sh -c 'echo hello-from-alfa && sleep 60'\n",
		"",
	)
	time.Sleep(200 * time.Millisecond)
	log := s.Send("log", "alfa")
	assert.True(t, strings.Contains(log, "hello-from-alfa"))
}

func TestLogUnknownProcess(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	result := s.Send("log", "nonexistent")
	assert.True(t, strings.HasPrefix(result, "error:"))
}
