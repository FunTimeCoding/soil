package server

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/integration/tester"
	"testing"
	"time"
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

func TestReloadProcfileUnchangedLeavesProcessRunning(t *testing.T) {
	marker, entry := countingEntry(t, "alfa")
	s := tester.New(t, entry, "")
	s.WaitOutput(t, "*alfa", "status")
	assert.Integer(t, 1, launchCount(t, marker))
	assert.String(t, "ok", s.Send("reload-procfile"))
	s.WaitOutput(t, "*alfa", "status")
	time.Sleep(300 * time.Millisecond)
	assert.Integer(t, 1, launchCount(t, marker))
}

func TestReloadProcfileRemovingOneLeavesOthersRunning(t *testing.T) {
	marker, entry := countingEntry(t, "alfa")
	s := tester.New(t, fmt.Sprintf("%sbravo: sleep 60\n", entry), "")
	s.WaitOutput(t, "*alfa\n*bravo", "status")
	s.WriteProcfile(entry)
	assert.String(t, "ok", s.Send("reload-procfile"))
	s.WaitOutput(t, "*alfa", "status")
	time.Sleep(300 * time.Millisecond)
	assert.Integer(t, 1, launchCount(t, marker))
}
