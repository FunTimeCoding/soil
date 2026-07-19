package server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/integration_test/tester"
	"testing"
)

func TestRestartProcess(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	result := s.Send("restart", "alfa")
	assert.String(t, "ok", result)
	s.WaitOutput(t, "*alfa", "status")
}

func TestStopProcess(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	result := s.Send("stop", "alfa")
	assert.String(t, "ok", result)
	s.WaitOutput(t, "alfa", "status")
}

func TestStartStoppedProcess(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	s.Send("stop", "alfa")
	s.WaitOutput(t, "alfa", "status")
	result := s.Send("start", "alfa")
	assert.String(t, "ok", result)
	s.WaitOutput(t, "*alfa", "status")
}
