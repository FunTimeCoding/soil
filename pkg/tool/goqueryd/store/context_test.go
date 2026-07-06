//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"testing"
)

func TestContextHierarchicalResolution(t *testing.T) {
	s, _ := indexedTestStore(t)
	defer s.Close()
	s.AddContext("test", separator.Slash, "root context")
	s.AddContext("test", "/tools/", "tools context")
	root := s.resolveContext("test", "alpha.md")
	assert.String(t, "root context", root)
	sub := s.resolveContext("test", "tools/gamma.md")
	assert.StringContains(t, "root context", sub)
	assert.StringContains(t, "tools context", sub)
}

func TestContextAttachedToSearchResults(t *testing.T) {
	s, _ := indexedTestStore(t)
	defer s.Close()
	s.AddContext("test", separator.Slash, "all documents")
	results := s.MustSearchKeyword("hybrid search pipeline", 10, "", false)
	assert.Count(t, 1, results)
	assert.String(t, "all documents", results[0].Context)
}

func TestContextAddOverwrites(t *testing.T) {
	s, _ := openTestStore(t)
	defer s.Close()
	directory := t.TempDir()
	s.AddCollection("test", directory, constant.DefaultGlob)
	s.AddContext("test", separator.Slash, "first")
	s.AddContext("test", separator.Slash, "second")
	entries := s.ListContexts()
	assert.Count(t, 1, entries)
	assert.String(t, "second", entries[0].Description)
}

func TestContextRemove(t *testing.T) {
	s, _ := openTestStore(t)
	defer s.Close()
	directory := t.TempDir()
	s.AddCollection("test", directory, constant.DefaultGlob)
	s.AddContext("test", separator.Slash, "to remove")
	removed := s.RemoveContext("test", separator.Slash)
	assert.True(t, removed)
	entries := s.ListContexts()
	assert.Count(t, 0, entries)
}

func TestContextRemoveNotFound(t *testing.T) {
	s, _ := openTestStore(t)
	defer s.Close()
	removed := s.RemoveContext("nonexistent", separator.Slash)
	assert.False(t, removed)
}
