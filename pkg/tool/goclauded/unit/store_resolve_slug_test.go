package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	"testing"
)

func TestResolveBySlug(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	s.Store.UpdateFields(
		"session-1",
		map[string]any{constant.Slug: "quiet-folding-heron"},
	)
	r, e := s.Store.ResolveSessionIdentifier("quiet-folding-heron")
	assert.FatalOnError(t, e)
	assert.True(t, r.Found())
	assert.String(t, "session-1", r.Identifier())
	assert.String(t, "slug", r.Matches[0].Field)
}

func TestResolveSharedSlugIsAmbiguous(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	s.EnsureSession("session-2")
	s.Store.UpdateFields(
		"session-1",
		map[string]any{constant.Slug: "quiet-folding-heron"},
	)
	s.Store.UpdateFields(
		"session-2",
		map[string]any{constant.Slug: "quiet-folding-heron"},
	)
	r, e := s.Store.ResolveSessionIdentifier("quiet-folding-heron")
	assert.FatalOnError(t, e)
	assert.True(t, r.Ambiguous())
	assert.Count(t, 2, r.Matches)
}

func TestResolveSlugEqualToOwnAliasMatchesOnce(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	s.Store.UpdateFields(
		"session-1",
		map[string]any{constant.Slug: "quiet-folding-heron"},
	)
	s.EditAlias("session-1", "quiet-folding-heron")
	r, e := s.Store.ResolveSessionIdentifier("quiet-folding-heron")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, r.Matches)
	assert.String(t, "alias", r.Matches[0].Field)
}
