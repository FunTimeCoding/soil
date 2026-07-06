//go:build local

package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
)

func TestComplete(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Store.Announce(r.Callsign, "fixing auth", "")
	topic := s.Store.CompleteTask(r.Callsign)
	s.Complete("session-1", r.Callsign, topic, "auth bug resolved")
	assert.String(
		t,
		"",
		s.Store.GetSession("session-1").Topic,
	)
	events := s.Store.Events(event_query.New().Kind(constant.Complete).SetLimit(10))
	assert.Count(t, 1, events)
	assert.String(
		t,
		"auth bug resolved",
		events[0].Metadata[constant.Message],
	)
	completions := s.Store.CompletionsBySession("session-1")
	assert.Count(t, 1, completions)
	assert.String(t, "auth bug resolved", completions[0].Summary)
}
