package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
	"testing"
)

func refusalFor(
	t *testing.T,
	s *service_tester.Tester,
	identifier string,
) string {
	t.Helper()
	r := s.Store.GetSession(identifier)
	assert.True(t, r != nil)
	result, e := s.Service.EmptyRefusal(r)
	assert.FatalOnError(t, e)
	assert.True(t, result != nil)

	return result.Error()
}

func TestEmptyRefusalAcceptsAHusk(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("husk")
	r := s.Store.GetSession("husk")
	result, e := s.Service.EmptyRefusal(r)
	assert.FatalOnError(t, e)
	assert.True(t, result == nil)
}

func TestEmptyRefusalNamesTurns(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("ghost")
	s.Store.Store.UpdateFields("ghost", map[string]any{"turn_count": 12})
	assert.String(t, "session has turns", refusalFor(t, s, "ghost"))
}

func TestEmptyRefusalNamesTranscriptLines(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("ghost")
	s.Store.Store.UpdateFields("ghost", map[string]any{"lines": 400})
	assert.String(t, "session has transcript lines", refusalFor(t, s, "ghost"))
}

func TestEmptyRefusalNamesCompletions(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("held")
	_, e := s.Store.Store.UpsertCompletion(
		"held",
		r.Callsign,
		constant.Complete,
		"a topic",
		"a summary",
	)
	assert.FatalOnError(t, e)
	assert.String(t, "session has completions", refusalFor(t, s, "held"))
}

func TestEmptyRefusalNamesSummary(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("held")
	assert.FatalOnError(
		t,
		s.Store.Store.UpsertSummary("held", r.Callsign, "a summary"),
	)
	assert.String(t, "session has a summary", refusalFor(t, s, "held"))
}

func TestEmptyRefusalNamesLabels(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("held")
	_, e := s.Store.Store.SetLabel("held", "role", "reviewer")
	assert.FatalOnError(t, e)
	assert.String(t, "session has labels", refusalFor(t, s, "held"))
}

func TestEmptyRefusalNamesPulses(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("held")
	assert.FatalOnError(
		t,
		s.Store.Store.SendPulse("held", r.Callsign, "a pulse"),
	)
	assert.String(t, "session has pulses", refusalFor(t, s, "held"))
}

func TestEmptyRefusalNamesContextLoads(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("held")
	load := context_load.New()
	load.SessionIdentifier = "held"
	load.CallIdentifier = "call-1"
	load.Reference = "memory-1"
	load.Kind = "memory"
	assert.FatalOnError(
		t,
		s.Store.Store.SaveContextLoads([]context_load.Load{*load}),
	)
	assert.String(t, "session has context loads", refusalFor(t, s, "held"))
}

func TestEmptyRefusalNamesEventsBeyondLifecycle(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("held")
	assert.FatalOnError(
		t,
		s.Store.Store.LogEvent(
			"held",
			constant.Announce,
			r.Callsign,
			map[string]string{constant.Topic: "a topic"},
		),
	)
	assert.String(
		t,
		"session has events beyond its lifecycle",
		refusalFor(t, s, "held"),
	)
}

func TestEmptyRefusalIgnoresLifecycleEvents(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("husk")
	assert.FatalOnError(t, s.Service.CloseSession("husk", "logout"))
	r := s.Store.GetSession("husk")
	result, e := s.Service.EmptyRefusal(r)
	assert.FatalOnError(t, e)
	assert.True(t, result == nil)
}
