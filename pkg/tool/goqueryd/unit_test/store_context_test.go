package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	strings "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"testing"
)

func TestContextHierarchicalResolution(t *testing.T) {
	s := indexedTestStore(t)
	defer s.Close()
	s.AddContext("test", strings.Slash, "root context")
	s.AddContext("test", "/tools/", "tools context")
	root := s.ResolveContext("test", "alpha.md")
	assert.String(t, "root context", root)
	sub := s.ResolveContext("test", "tools/gamma.md")
	assert.StringContains(t, "root context", sub)
	assert.StringContains(t, "tools context", sub)
}

func TestContextAttachedToSearchResults(t *testing.T) {
	s := indexedTestStore(t)
	defer s.Close()
	s.AddContext("test", strings.Slash, "all documents")
	results := s.MustSearchKeyword(
		"hybrid search pipeline",
		10,
		"",
		false,
		nil,
	)
	assert.Count(t, 1, results)
	assert.String(t, "all documents", results[0].Context)
}

func TestContextAddOverwrites(t *testing.T) {
	s := openTestStore(t)
	defer s.Close()
	directory := t.TempDir()
	s.AddCollection("test", directory, constant.DefaultGlob)
	s.AddContext("test", strings.Slash, "first")
	s.AddContext("test", strings.Slash, "second")
	entries := s.ListContexts()
	assert.Count(t, 1, entries)
	assert.String(t, "second", entries[0].Description)
}

func TestContextRemove(t *testing.T) {
	s := openTestStore(t)
	defer s.Close()
	directory := t.TempDir()
	s.AddCollection("test", directory, constant.DefaultGlob)
	s.AddContext("test", strings.Slash, "to remove")
	removed := s.RemoveContext("test", strings.Slash)
	assert.True(t, removed)
	entries := s.ListContexts()
	assert.Count(t, 0, entries)
}

func TestContextRemoveNotFound(t *testing.T) {
	s := openTestStore(t)
	defer s.Close()
	removed := s.RemoveContext("nonexistent", strings.Slash)
	assert.False(t, removed)
}
