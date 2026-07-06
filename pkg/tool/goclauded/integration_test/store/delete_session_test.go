//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/store_tester"
	"testing"
)

func TestDeleteSession(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	assert.True(t, s.GetSession("session-1") != nil)
	assert.FatalOnError(t, s.Store.DeleteSession("session-1"))
	assert.True(t, s.GetSession("session-1") == nil)
}

func TestDeleteSessionCascadesChildRecords(t *testing.T) {
	s := store_tester.New(t)
	r := s.EnsureSession("session-1")
	assert.FatalOnError(t, s.Store.Announce(r.Callsign, "test topic", ""))
	assert.FatalOnError(
		t,
		s.Store.LogEvent(
			"session-1",
			constant.Announce,
			"Ash",
			map[string]string{constant.Topic: "test topic"},
		),
	)
	_, e := s.Store.UpsertCompletion(
		"session-1",
		"Ash",
		constant.Complete,
		"test topic",
		"test summary",
	)
	assert.FatalOnError(t, e)
	assert.FatalOnError(
		t,
		s.Store.UpsertSummary("session-1", "Ash", "test summary body"),
	)
	_, se := s.Store.SetLabel("session-1", "role", "reviewer")
	assert.FatalOnError(t, se)
	assert.FatalOnError(
		t,
		s.Store.SendPulse("session-1", "Ash", "test pulse"),
	)
	completions := s.CompletionsBySession("session-1")
	assert.Count(t, 1, completions)
	summary := s.SummaryBySession("session-1")
	assert.True(t, summary != "")
	labels, le := s.Store.LabelsBySession("session-1")
	assert.FatalOnError(t, le)
	assert.Count(t, 1, labels)
	pulses, pe := s.Store.PulsesBySession("session-1")
	assert.FatalOnError(t, pe)
	assert.Count(t, 1, pulses)
	eventCount, ce := s.Store.CountEvents()
	assert.FatalOnError(t, ce)
	assert.True(t, eventCount > 0)
	assert.FatalOnError(t, s.Store.DeleteSession("session-1"))
	assert.True(t, s.GetSession("session-1") == nil)
	completions = s.CompletionsBySession("session-1")
	assert.Count(t, 0, completions)
	summary = s.SummaryBySession("session-1")
	assert.String(t, "", summary)
	labels, le = s.Store.LabelsBySession("session-1")
	assert.FatalOnError(t, le)
	assert.Count(t, 0, labels)
	pulses, pe = s.Store.PulsesBySession("session-1")
	assert.FatalOnError(t, pe)
	assert.Count(t, 0, pulses)
	eventCount, ce = s.Store.CountEvents()
	assert.FatalOnError(t, ce)
	assert.Integer64(t, 0, eventCount)
}
