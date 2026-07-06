//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/store_tester"
	"testing"
)

func TestResolvePrefixMatch(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("abc-1234")
	r, e := s.Store.ResolveSessionIdentifier("abc")
	assert.FatalOnError(t, e)
	assert.True(t, r.Found())
	assert.String(t, "abc-1234", r.Identifier())
}

func TestResolvePrefixAmbiguous(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("abc-1111")
	s.EnsureSession("abc-2222")
	r, e := s.Store.ResolveSessionIdentifier("abc")
	assert.FatalOnError(t, e)
	assert.True(t, r.Ambiguous())
	assert.Count(t, 2, r.Matches)
}

func TestResolvePrefixNoMatch(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("abc-1234")
	r, e := s.Store.ResolveSessionIdentifier("xyz")
	assert.FatalOnError(t, e)
	assert.False(t, r.Found())
	assert.False(t, r.Ambiguous())
}

func TestResolveDedup(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("abc-1234")
	r, e := s.Store.ResolveSessionIdentifier("abc-1234")
	assert.FatalOnError(t, e)
	assert.True(t, r.Found())
	assert.Count(t, 1, r.Matches)
}

func TestResolveByCallsign(t *testing.T) {
	s := store_tester.New(t)
	r := s.EnsureSession("session-1")
	result, e := s.Store.ResolveSessionIdentifier(r.Callsign)
	assert.FatalOnError(t, e)
	assert.True(t, result.Found())
	assert.String(t, "session-1", result.Identifier())
}
