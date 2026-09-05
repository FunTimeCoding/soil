package server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/integration/tester"
	"strings"
	"testing"
)

func TestLogShowsCurrentGenerationAfterRestart(t *testing.T) {
	s := tester.New(t, "alfa: sh -c 'echo mark && sleep 60'\n", "")
	s.WaitContains(t, "mark", "log", "alfa")
	s.Send("restart", "alfa")
	s.WaitContains(t, "older lines", "log", "alfa")
	s.WaitContains(t, "mark", "log", "alfa")
	result := s.Send("log", "alfa")
	assert.Integer(t, 1, strings.Count(result, "mark"))
	assert.True(t, strings.Contains(result, "(2 older lines"))
	assert.StringNotContains(t, "Terminating", result)
}

func TestLogAllShowsEveryGeneration(t *testing.T) {
	s := tester.New(t, "alfa: sh -c 'echo mark && sleep 60'\n", "")
	s.WaitContains(t, "mark", "log", "alfa")
	s.Send("restart", "alfa")
	s.WaitContains(t, "older lines", "log", "alfa")
	s.WaitContains(t, "mark", "log", "alfa")
	result := s.Send("log", "alfa", "all")
	assert.Integer(t, 2, strings.Count(result, "mark"))
	assert.True(t, strings.Contains(result, "Terminating alfa"))
}

func TestLogClearDiscardsHistory(t *testing.T) {
	s := tester.New(t, "alfa: sh -c 'echo mark && sleep 60'\n", "")
	s.WaitContains(t, "mark", "log", "alfa")
	assert.String(t, "ok", s.Send("log", "alfa", "clear"))
	assert.String(t, "ok", s.Send("log", "alfa", "all"))
}

func TestLogUnknownOption(t *testing.T) {
	s := tester.New(t, "alfa: sleep 60\n", "")
	assert.True(t, strings.HasPrefix(s.Send("log", "alfa", "wipe"), "error:"))
}

func TestReloadProcfileKeepsLogHistory(t *testing.T) {
	s := tester.New(t, "alfa: sh -c 'echo mark && sleep 60'\n", "")
	s.WaitContains(t, "mark", "log", "alfa")
	s.WriteProcfile("alfa: sh -c 'echo mark && sleep 120'\n")
	assert.String(t, "ok", s.Send("reload-procfile"))
	s.WaitContains(t, "older lines", "log", "alfa")
	s.WaitContains(t, "mark", "log", "alfa")
	assert.Integer(t, 2, strings.Count(s.Send("log", "alfa", "all"), "mark"))
}
