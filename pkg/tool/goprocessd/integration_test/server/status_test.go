package server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/integration_test/tester"
	"strings"
	"testing"
)

func TestStatusShowsRunning(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\nbravo: sleep 60\n", "")
	result := s.Send("status")
	lines := strings.Split(result, "\n")
	assert.Integer(t, 2, len(lines))
	assert.String(t, "*alfa", lines[0])
	assert.String(t, "*bravo", lines[1])
}

func TestListShowsAllProcesses(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\nbravo: sleep 60\n", "")
	result := s.Send("list")
	lines := strings.Split(result, "\n")
	assert.Integer(t, 2, len(lines))
	assert.String(t, "alfa", lines[0])
	assert.String(t, "bravo", lines[1])
}
