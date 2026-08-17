package service

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/claude/session"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
	"os"
	"path/filepath"
	"testing"
)

func writeContextLoadFile(
	t *testing.T,
	harbor string,
	identifier string,
) {
	t.Helper()
	errors.PanicOnError(
		os.WriteFile(
			filepath.Join(
				harbor,
				join.Empty(identifier, constant.NotationLogExtension),
			),
			[]byte(fixture.Read("claude", "context-loads.jsonl")),
			0o644,
		),
	)
}

func loadAt(
	t *testing.T,
	loads []context_load.Load,
	index int,
) *context_load.Load {
	t.Helper()

	if index >= len(loads) {
		t.Fatalf("expected at least %d loads, got %d", index+1, len(loads))
	}

	return &loads[index]
}

func TestContextLoadsRecordMemoriesAndModes(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("context-session")
	writeContextLoadFile(t, s.Harbor, "context-session")
	s.Service.EnrichSession("context-session")
	loads, e := s.Service.ContextLoadsBySession("context-session")
	assert.FatalOnError(t, e)
	assert.Integer(t, 8, len(loads))
	mode := loadAt(t, loads, 0)
	assert.String(t, "mode", mode.Kind)
	assert.String(t, "build", mode.Reference)
	always := loadAt(t, loads, 1)
	assert.String(t, "memory", always.Kind)
	assert.String(t, "102", always.Reference)
	assert.String(t, "bravo", always.Name)
	assert.String(t, "always", always.Tier)
	relevant := loadAt(t, loads, 3)
	assert.String(t, "104", relevant.Reference)
	assert.String(t, "relevant", relevant.Tier)
	fetched := loadAt(t, loads, 4)
	assert.String(t, "memory", fetched.Kind)
	assert.String(t, "105", fetched.Reference)
	assert.String(t, "golf", fetched.Name)
	assert.String(t, "", fetched.Tier)
	found := loadAt(t, loads, 6)
	assert.String(t, "search", found.Kind)
	assert.String(t, "107", found.Reference)
	assert.String(t, "example query", found.Query)
}

func TestContextLoadsSkipIndexTier(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("index-session")
	writeContextLoadFile(t, s.Harbor, "index-session")
	s.Service.EnrichSession("index-session")
	loads, e := s.Service.ContextLoadsBySession("index-session")
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
	writeContextLoadFile(t, s.Harbor, "replay-session")
	s.Service.EnrichSession("replay-session")
	first, e := s.Service.ContextLoadsBySession("replay-session")
	assert.FatalOnError(t, e)
	assert.Integer(t, 8, len(first))
	result := s.Service.ColdBackfillAllSessions()
	assert.Integer(t, 1, result.Enriched)
	second, f := s.Service.ContextLoadsBySession("replay-session")
	assert.FatalOnError(t, f)
	assert.Integer(t, len(first), len(second))
}
