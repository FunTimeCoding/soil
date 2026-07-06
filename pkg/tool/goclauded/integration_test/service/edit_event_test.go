//go:build local

package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/event_query"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"testing"
)

func TestEditEventSummaryCascade(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("session-1")
	assert.FatalOnError(
		t,
		s.Service.Summarize("session-1", "Cedar", "original summary"),
	)
	events := s.Store.Events(event_query.New().Kind(constant.Summarize).SetLimit(10))
	assert.Count(t, 1, events)
	_, e := s.Service.EditEvent(events[0].Identifier, "revised summary")
	assert.FatalOnError(t, e)
	body := s.Store.SummaryBySession("session-1")
	assert.String(t, "revised summary", body)
}

func TestEditEventCompletionCascade(t *testing.T) {
	s := service_tester.New(t)
	r := s.Store.EnsureSession("session-1")
	s.Store.Announce(r.Callsign, "fixing auth", "")
	topic := s.Store.CompleteTask(r.Callsign)
	s.Complete("session-1", r.Callsign, topic, "original message")
	events := s.Store.Events(event_query.New().Kind(constant.Complete).SetLimit(10))
	assert.Count(t, 1, events)
	_, f := s.Service.EditEvent(events[0].Identifier, "corrected message")
	assert.FatalOnError(t, f)
	completions := s.Store.CompletionsBySession("session-1")
	assert.Count(t, 1, completions)
	assert.String(t, "corrected message", completions[0].Summary)
}
