package server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/integration_test/tester"
	"strings"
	"testing"
	"time"
)

func TestRestartProcess(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	result := s.Send("restart", "alfa")
	assert.String(t, "ok", result)
	time.Sleep(200 * time.Millisecond)
	status := s.Send("status")
	assert.String(t, "*alfa", strings.TrimSpace(status))
}

func TestStopProcess(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	result := s.Send("stop", "alfa")
	assert.String(t, "ok", result)
	time.Sleep(200 * time.Millisecond)
	status := s.Send("status")
	assert.String(t, " alfa", status)
}

func TestStartStoppedProcess(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	s.Send("stop", "alfa")
	time.Sleep(200 * time.Millisecond)
	result := s.Send("start", "alfa")
	assert.String(t, "ok", result)
	time.Sleep(200 * time.Millisecond)
	status := s.Send("status")
	assert.String(t, "*alfa", strings.TrimSpace(status))
}
