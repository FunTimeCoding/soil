//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"testing"
)

func TestSearchKeywordFindsDocument(t *testing.T) {
	s, _ := indexedTestStore(t)
	defer s.Close()
	results := s.MustSearchKeyword("hybrid search pipeline", 10, "", false)
	assert.Count(t, 1, results)
	assert.String(t, "Search Pipeline", results[0].Title)
}

func TestSearchKeywordRanking(t *testing.T) {
	s, _ := indexedTestStore(t)
	defer s.Close()
	results := s.MustSearchKeyword("documents", 10, "", false)
	assert.Greater(t, 0, results[0].Score)
	assert.Greater(t, 1, float64(len(results)))
}

func TestSearchKeywordCollectionFilter(t *testing.T) {
	s, _ := openTestStore(t)
	defer s.Close()
	first := t.TempDir()
	second := t.TempDir()
	writeFixture(t, first, "one.md", "# One\n\nUnique keyword aardvark.\n")
	writeFixture(t, second, "two.md", "# Two\n\nUnique keyword aardvark.\n")
	s.AddCollection("first", first, constant.DefaultGlob)
	s.AddCollection("second", second, constant.DefaultGlob)
	s.Index("first")
	s.Index("second")
	all := s.MustSearchKeyword("aardvark", 10, "", false)
	assert.Count(t, 2, all)
	filtered := s.MustSearchKeyword("aardvark", 10, "first", false)
	assert.Count(t, 1, filtered)
	assert.String(t, "first", filtered[0].Collection)
}

func TestSearchKeywordVirtualPath(t *testing.T) {
	s, _ := indexedTestStore(t)
	defer s.Close()
	results := s.MustSearchKeyword("hybrid search pipeline", 10, "", false)
	assert.String(t, "qmd://test/alpha.md", results[0].VirtualPath)
}

func TestSearchKeywordNegation(t *testing.T) {
	s, _ := indexedTestStore(t)
	defer s.Close()
	all := s.MustSearchKeyword("search documents", 10, "", false)
	negated := s.MustSearchKeyword("search -pipeline", 10, "", false)
	assert.Greater(t, float64(len(negated)), float64(len(all)))
}

func TestSearchKeywordQuotedPhrase(t *testing.T) {
	s, _ := openTestStore(t)
	defer s.Close()
	directory := t.TempDir()
	writeFixture(
		t,
		directory,
		"exact.md",
		"# Exact\n\nThe quick brown fox jumps over the lazy dog.\n",
	)
	writeFixture(
		t,
		directory,
		"partial.md",
		"# Partial\n\nThe quick fox is brown and lazy.\n",
	)
	s.AddCollection("test", directory, constant.DefaultGlob)
	s.Index("test")
	results := s.MustSearchKeyword(`"quick brown fox"`, 10, "", false)
	assert.Count(t, 1, results)
	assert.String(t, "Exact", results[0].Title)
}
