//go:build local

package store

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/service/argument/edit_session"
	"testing"
)

func edit(
	s *store_tester.Tester,
	identifier string,
	a *edit_session.Session,
) {
	s.EditSession(identifier, a)
}

func editAlias(
	s *store_tester.Tester,
	identifier string,
	alias string,
) {
	edit(s, identifier, edit_session.New().WithAlias(alias))
}

func editDescription(
	s *store_tester.Tester,
	identifier string,
	description string,
) {
	edit(s, identifier, edit_session.New().WithDescription(description))
}

func TestSetAndGetAlias(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	editAlias(s, "session-1", "my-project")
	assert.String(
		t,
		"my-project",
		s.GetSession("session-1").AliasValue(),
	)
}

func TestSetAliasOverwrite(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	editAlias(s, "session-1", "first-name")
	editAlias(s, "session-1", "second-name")
	assert.String(
		t,
		"second-name",
		s.GetSession("session-1").AliasValue(),
	)
}

func TestSetAndGetDescription(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	editDescription(s, "session-1", "Fixed the auth bug")
	assert.String(
		t,
		"Fixed the auth bug",
		s.GetSession("session-1").Description,
	)
}

func TestSetDescriptionWithoutAlias(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	editDescription(s, "session-1", "Standalone description")
	e := s.GetSession("session-1")
	assert.String(t, "", e.AliasValue())
	assert.String(t, "Standalone description", e.Description)
}

func TestSetAliasAndDescriptionIndependently(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	editAlias(s, "session-1", "my-project")
	editDescription(s, "session-1", "Refactored the CLI")
	e := s.GetSession("session-1")
	assert.String(t, "my-project", e.AliasValue())
	assert.String(t, "Refactored the CLI", e.Description)
}

func TestSetBothAtOnce(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	edit(
		s,
		"session-1",
		edit_session.New().
			WithAlias("my-project").
			WithDescription("Refactored the CLI"),
	)
	e := s.GetSession("session-1")
	assert.String(t, "my-project", e.AliasValue())
	assert.String(t, "Refactored the CLI", e.Description)
}

func TestSetDescriptionDoesNotClearAlias(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	editAlias(s, "session-1", "my-project")
	editDescription(s, "session-1", "New description")
	assert.String(
		t,
		"my-project",
		s.GetSession("session-1").AliasValue(),
	)
}

func TestSetAliasDoesNotClearDescription(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	editDescription(s, "session-1", "Important work")
	editAlias(s, "session-1", "renamed")
	assert.String(
		t,
		"Important work",
		s.GetSession("session-1").Description,
	)
}

func TestEditTopic(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	edit(s, "session-1", edit_session.New().WithTopic("debugging auth"))
	assert.String(
		t,
		"debugging auth",
		s.GetSession("session-1").Topic,
	)
}

func TestNoOpEdit(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	editAlias(s, "session-1", constant.FixtureBefore)
	edit(s, "session-1", edit_session.New())
	assert.String(
		t,
		"before",
		s.GetSession("session-1").AliasValue(),
	)
}

func TestResolveByAlias(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	editAlias(s, "session-1", "my-project")
	resolved, e := s.Store.ResolveSessionIdentifier("my-project")
	assert.FatalOnError(t, e)
	assert.True(t, resolved.Found())
	assert.String(t, "session-1", resolved.Identifier())
}

func TestResolveByIdentifier(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	resolved, e := s.Store.ResolveSessionIdentifier("session-1")
	assert.FatalOnError(t, e)
	assert.True(t, resolved.Found())
	assert.String(t, "session-1", resolved.Identifier())
}

func TestAliasedSessions(t *testing.T) {
	s := store_tester.New(t)
	s.EnsureSession("session-1")
	s.EnsureSession("session-2")
	s.EnsureSession("session-3")
	editAlias(s, "session-1", "first")
	editAlias(s, "session-2", "second")
	aliased := s.Store.AliasedSessions()
	assert.Count(t, 2, aliased)
}
