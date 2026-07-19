package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
)

func TestUpdate(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Store.Announce(r.Callsign, "initial topic", "")
	s.Update(
		"session-1",
		r.Callsign,
		"new scope",
		"completed initial work",
		"pkg/tool",
	)
	e := s.Store.GetSession("session-1")
	assert.String(t, "new scope", e.Topic)
	assert.String(t, "pkg/tool", e.Files)
	events := s.Store.Events(event_query.New().Kind(constant.Update).SetLimit(10))
	assert.Count(t, 1, events)
	completions := s.Store.CompletionsBySession("session-1")
	assert.Count(t, 1, completions)
	assert.String(t, "new scope", completions[0].Topic)
	assert.String(t, "completed initial work", completions[0].Summary)
}
