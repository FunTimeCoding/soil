package server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/integration/tester"
	"testing"
	"time"
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

func TestRestartAllLeavesEveryProcessRunning(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\nbravo: sleep 60\n", "")
	s.WaitOutput(t, "*alfa\n*bravo", "status")
	assert.String(t, "ok", s.Send("restart-all"))
	s.WaitOutput(t, "*alfa\n*bravo", "status")
	time.Sleep(time.Second)
	assert.String(t, "*alfa\n*bravo", s.Send("status"))
}

func TestRestartingEveryProcessOneByOneLeavesThemRunning(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\nbravo: sleep 60\n", "")
	s.WaitOutput(t, "*alfa\n*bravo", "status")
	assert.String(t, "ok", s.Send("restart", "alfa"))
	s.WaitOutput(t, "*alfa\n*bravo", "status")
	assert.String(t, "ok", s.Send("restart", "bravo"))
	s.WaitOutput(t, "*alfa\n*bravo", "status")
	time.Sleep(time.Second)
	assert.String(t, "*alfa\n*bravo", s.Send("status"))
}
