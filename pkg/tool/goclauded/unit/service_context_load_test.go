package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/fixture"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"testing"
)

func TestContextLoadsRecordMemoriesAndModes(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("context-session")
	s.WriteContextLoadFile("context-session")
	s.Transcript.EnrichSession("context-session")
	loads, e := s.Transcript.SessionContextLoads(t.Context(), "context-session")
	assert.FatalOnError(t, e)
	assert.Integer(t, 11, len(loads))
	mode := fixture.LoadAt(t, loads, 0)
	assert.String(t, "mode", mode.Kind)
	assert.String(t, "build", mode.Reference)
	always := fixture.LoadAt(t, loads, 1)
	assert.String(t, "memory", always.Kind)
	assert.String(t, "102", always.Reference)
	assert.String(t, "bravo", always.Name)
	assert.String(t, "always", always.Tier)
	relevant := fixture.LoadAt(t, loads, 3)
	assert.String(t, "104", relevant.Reference)
	assert.String(t, "relevant", relevant.Tier)
	fetched := fixture.LoadAt(t, loads, 4)
	assert.String(t, "memory", fetched.Kind)
	assert.String(t, "105", fetched.Reference)
	assert.String(t, "golf", fetched.Name)
	assert.String(t, "", fetched.Tier)
	found := fixture.LoadAt(t, loads, 6)
	assert.String(t, "search", found.Kind)
	assert.String(t, "107", found.Reference)
	assert.String(t, "example query", found.Query)
	parent := fixture.LoadAt(t, loads, 8)
	assert.String(t, "memory", parent.Kind)
	assert.String(t, "109", parent.Reference)
	assert.String(t, "kilo", parent.Name)
	child := fixture.LoadAt(t, loads, 10)
	assert.String(t, "111", child.Reference)
	assert.String(t, "mike", child.Name)
}

func TestContextLoadsSkipIndexTier(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("index-session")
	s.WriteContextLoadFile("index-session")
	s.Transcript.EnrichSession("index-session")
	loads, e := s.Transcript.SessionContextLoads(t.Context(), "index-session")
	assert.FatalOnError(t, e)

	for _, entry := range loads {
		assert.True(t, entry.Reference != "201")
		assert.True(t, entry.Reference != "202")
	}
}

func TestContextLoadsSurviveColdReplay(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("replay-session")
	s.Client.AddSession(session.New("replay-session"))
	s.WriteContextLoadFile("replay-session")
	s.Transcript.EnrichSession("replay-session")
	first, e := s.Transcript.SessionContextLoads(t.Context(), "replay-session")
	assert.FatalOnError(t, e)
	assert.Integer(t, 11, len(first))
	result := s.Transcript.ColdBackfillAllSessions()
	assert.Integer(t, 1, result.Enriched)
	second, f := s.Transcript.SessionContextLoads(t.Context(), "replay-session")
	assert.FatalOnError(t, f)
	assert.Integer(t, len(first), len(second))
}

func TestContextLoadsRecordTextProfileTiers(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("text-session")
	s.WriteTranscript("text-session", "context-loads-text.jsonl")
	s.Transcript.EnrichSession("text-session")
	loads, e := s.Transcript.SessionContextLoads(t.Context(), "text-session")
	assert.FatalOnError(t, e)
	assert.Integer(t, 11, len(loads))
	always := fixture.LoadAt(t, loads, 1)
	assert.String(t, "102", always.Reference)
	assert.String(t, "bravo", always.Name)
	assert.String(t, "always", always.Tier)
	second := fixture.LoadAt(t, loads, 2)
	assert.String(t, "103", second.Reference)
	assert.String(t, "always", second.Tier)
	relevant := fixture.LoadAt(t, loads, 3)
	assert.String(t, "104", relevant.Reference)
	assert.String(t, "delta", relevant.Name)
	assert.String(t, "relevant", relevant.Tier)

	for _, entry := range loads {
		assert.True(t, entry.Reference != "201")
	}

	assert.Count(t, 0, s.Reporter.Events())
}

func TestContextLoadsFollowWholeReadOfPersistedProfile(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("persisted-session")
	s.WriteTranscript("persisted-session", "context-loads-persisted.jsonl")
	s.Transcript.EnrichSession("persisted-session")
	loads, e := s.Transcript.SessionContextLoads(
		t.Context(),
		"persisted-session",
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 11, len(loads))
	always := fixture.LoadAt(t, loads, 1)
	assert.String(t, "102", always.Reference)
	assert.String(t, "always", always.Tier)
	relevant := fixture.LoadAt(t, loads, 3)
	assert.String(t, "104", relevant.Reference)
	assert.String(t, "relevant", relevant.Tier)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestContextLoadsReportRecentPersistedProfileOnce(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("persisted-session")
	s.Store.EnsureSession("second-persisted-session")
	s.WriteRecentTranscript(
		"persisted-session",
		"context-loads-persisted.jsonl",
	)
	s.WriteRecentTranscript(
		"second-persisted-session",
		"context-loads-persisted.jsonl",
	)
	s.Transcript.EnrichSession("persisted-session")
	s.Transcript.EnrichSession("second-persisted-session")
	assert.Count(t, 1, s.Reporter.Events())
}

func TestContextLoadsReportUnreadableProfileOnce(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("drift-session")
	s.Store.EnsureSession("second-drift-session")
	s.WriteTranscript("drift-session", "context-loads-drift.jsonl")
	s.WriteTranscript("second-drift-session", "context-loads-drift.jsonl")
	s.Transcript.EnrichSession("drift-session")
	s.Transcript.EnrichSession("second-drift-session")
	assert.Count(t, 1, s.Reporter.Events())
}
