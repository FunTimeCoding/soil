//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestSetAndGetSourceType(t *testing.T) {
	s, _ := openTestStore(t)
	defer s.Close()
	s.SetSourceType("docs", "design/", "design-doc")
	result := s.GetSourceType("docs", "design/")
	assert.String(t, "design-doc", result)
}

func TestGetSourceTypeEmpty(t *testing.T) {
	s, _ := openTestStore(t)
	defer s.Close()
	result := s.GetSourceType("docs", "nonexistent/")
	assert.String(t, "", result)
}

func TestRemoveSourceType(t *testing.T) {
	s, _ := openTestStore(t)
	defer s.Close()
	s.SetSourceType("docs", "temp/", "scratch")
	s.SetSourceType("docs", "temp/", "")
	result := s.GetSourceType("docs", "temp/")
	assert.String(t, "", result)
}

func TestResolveSourceTypeMostSpecificWins(t *testing.T) {
	s, _ := openTestStore(t)
	defer s.Close()
	s.SetSourceType("docs", "plan/", "plan")
	s.SetSourceType("docs", "plan/seed/", "seed")
	broad := s.ResolveSourceType("docs", "plan/overview.md")
	assert.String(t, "plan", broad)
	specific := s.ResolveSourceType("docs", "plan/seed/goqueryd.md")
	assert.String(t, "seed", specific)
}

func TestResolveSourceTypeGlobalFallback(t *testing.T) {
	s, _ := openTestStore(t)
	defer s.Close()
	s.SetSourceType("", "memory/", "memory")
	result := s.ResolveSourceType("any-collection", "memory/42.md")
	assert.String(t, "memory", result)
}

func TestResolveSourceTypeCollectionOverridesGlobal(t *testing.T) {
	s, _ := openTestStore(t)
	defer s.Close()
	s.SetSourceType("", "session/", "generic")
	s.SetSourceType("summaries", "session/", "session-summary")
	global := s.ResolveSourceType("other", "session/abc.md")
	assert.String(t, "generic", global)
	specific := s.ResolveSourceType("summaries", "session/abc.md")
	assert.String(t, "session-summary", specific)
}
