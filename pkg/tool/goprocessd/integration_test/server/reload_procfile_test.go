package server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/integration_test/tester"
	"testing"
)

func TestReloadProcfileAddsNewEntry(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	s.WriteProcfile("alfa: sleep 60\nbravo: sleep 60\n")
	result := s.Send("reload-procfile")
	assert.String(t, "ok", result)
	s.WaitOutput(t, "*alfa\n*bravo", "status")
}

func TestReloadProcfileRemovesEntry(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\nbravo: sleep 60\n", "")
	s.WriteProcfile("alfa: sleep 60\n")
	result := s.Send("reload-procfile")
	assert.String(t, "ok", result)
	s.WaitOutput(t, "alfa", "list")
}

func TestReloadProcfileChangedCommand(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	s.WriteProcfile("alfa: sleep 120\n")
	result := s.Send("reload-procfile")
	assert.String(t, "ok", result)
	s.WaitOutput(t, "*alfa", "status")
}
