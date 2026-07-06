//go:build local

package server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/integration_test/tester"
	"strings"
	"testing"
	"time"
)

func TestReloadProcfileAddsNewEntry(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	s.WriteProcfile("alfa: sleep 60\nbravo: sleep 60\n")
	result := s.Send("reload-procfile")
	assert.String(t, "ok", result)
	time.Sleep(200 * time.Millisecond)
	status := s.Send("status")
	lines := strings.Split(status, "\n")
	assert.Integer(t, 2, len(lines))
	assert.String(t, "*alfa", lines[0])
	assert.String(t, "*bravo", lines[1])
}

func TestReloadProcfileRemovesEntry(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\nbravo: sleep 60\n", "")
	s.WriteProcfile("alfa: sleep 60\n")
	result := s.Send("reload-procfile")
	assert.String(t, "ok", result)
	time.Sleep(200 * time.Millisecond)
	list := s.Send("list")
	assert.String(t, "alfa", strings.TrimSpace(list))
}

func TestReloadProcfileChangedCommand(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	s.WriteProcfile("alfa: sleep 120\n")
	result := s.Send("reload-procfile")
	assert.String(t, "ok", result)
	time.Sleep(500 * time.Millisecond)
	status := s.Send("status")
	assert.String(t, "*alfa", strings.TrimSpace(status))
}
