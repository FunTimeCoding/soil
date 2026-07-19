package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/store_tester"
	"testing"
)

func TestResolveByName(t *testing.T) {
	s := store_tester.New(t)
	r := s.EnsureSession("session-1")
	resolved, e := s.Store.ResolveByCallsign(r.Callsign)
	assert.FatalOnError(t, e)
	assert.String(t, "session-1", resolved)
}

func TestResolveByNameNotFound(t *testing.T) {
	s := store_tester.New(t)
	resolved, e := s.Store.ResolveByCallsign("nonexistent")
	assert.FatalOnError(t, e)
	assert.String(t, "", resolved)
}

func TestAliasOwner(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	editAlias(s, "session-1", "my-alias")
	owner, e := s.Store.AliasOwner("my-alias")
	assert.FatalOnError(t, e)
	assert.String(t, "session-1", owner)
}

func TestAliasOwnerNotFound(t *testing.T) {
	s := store_tester.New(t)
	owner, e := s.Store.AliasOwner("nonexistent")
	assert.FatalOnError(t, e)
	assert.String(t, "", owner)
}

func TestSessionByName(t *testing.T) {
	s := store_tester.New(t)
	r := s.EnsureSession("session-1")
	result, e := s.Store.SessionByCallsign(r.Callsign)
	assert.FatalOnError(t, e)
	assert.String(t, "session-1", result.Identifier)
}

func TestSessionByNameNotFound(t *testing.T) {
	s := store_tester.New(t)
	result, e := s.Store.SessionByCallsign("nonexistent")
	assert.FatalOnError(t, e)
	assert.True(t, result == nil)
}
