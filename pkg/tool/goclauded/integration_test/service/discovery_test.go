package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/service_tester"
	"path/filepath"
	"testing"
)

func TestCheckConsistencyDiscoversUnregistered(t *testing.T) {
	s := service_tester.New(t)
	writeSessionFile(s.Harbor, "jsonl-only", "test-session")
	s.Service.PopulateCache()
	s.Service.CheckConsistency()
	e := s.Store.GetSession("jsonl-only")
	assert.True(t, e != nil)
	assert.String(t, "test-session", e.Slug)
	assert.True(t, e.Lines == 2)
	assert.True(t, e.Callsign == nil)
	assert.String(t, "", e.Name)
}

func TestCheckConsistencySkipsExisting(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("registered")
	writeSessionFile(s.Harbor, "registered", "already-known")
	s.Service.PopulateCache()
	s.Service.CheckConsistency()
	e := s.Store.GetSession("registered")
	assert.True(t, e.Callsign != nil)
}

func TestBackfillSessionsEnriches(t *testing.T) {
	s := service_tester.New(t)
	s.Store.EnsureSession("session-1")
	writeSessionFile(s.Harbor, "session-1", "my-slug")
	s.Service.PopulateCache()
	s.Service.BackfillSessions()
	e := s.Store.GetSession("session-1")
	assert.String(t, "my-slug", e.Slug)
	assert.True(t, e.Lines == 2)
	assert.String(t, "/home/user", e.WorkDirectory)
	assert.String(t, "main", e.Branch)
	assert.String(t, "can you help me fix the login bug", e.FirstMessage)
	assert.True(t, e.TurnCount == 2)
}

func writeSessionFile(
	harbor string,
	identifier string,
	slug string,
) {
	path := filepath.Join(harbor, fmt.Sprintf("%s.jsonl", identifier))
	f := system.Create(path)
	_, e := fmt.Fprintf(
		f,
		"{\"type\":\"user\",\"timestamp\":\"2026-05-21T10:00:00Z\",\"sessionId\":\"%s\",\"slug\":\"%s\",\"cwd\":\"/home/user\",\"gitBranch\":\"main\",\"message\":{\"role\":\"user\",\"content\":\"can you help me fix the login bug\"}}\n{\"type\":\"user\",\"timestamp\":\"2026-05-21T10:02:00Z\",\"sessionId\":\"%s\",\"message\":{\"role\":\"user\",\"content\":\"now update the tests to match\"}}\n",
		identifier,
		slug,
		identifier,
	)
	errors.PanicOnError(e)
	errors.PanicClose(f)
}
