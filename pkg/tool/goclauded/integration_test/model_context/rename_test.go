//go:build local

package model_context

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration_test/base"
	"testing"
)

func TestEditSessionAliasSelf(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working")
	a.MustCallTool(
		constant.EditSession,
		map[string]any{
			constant.Alias: "my-project",
		},
	)
	roster := a.MustCallTool(constant.Roster, map[string]any{})
	assert.StringContains(t, "my-project", roster)
}

func TestEditSessionAliasOther(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	b := s.NewSession(t)
	defer b.Close()
	a.Announce(a.Name(), "editor")
	b.Announce(b.Name(), constant.FixtureTarget)
	a.MustCallTool(
		constant.EditSession,
		map[string]any{
			constant.Alias:  "the-target",
			constant.Target: b.Name(),
		},
	)
	roster := a.MustCallTool(constant.Roster, map[string]any{})
	assert.StringContains(t, "the-target", roster)
}

func TestEditSessionDescription(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working")
	a.MustCallTool(
		constant.EditSession,
		map[string]any{
			constant.Description: "Fixed the auth bug",
		},
	)
	assert.String(
		t,
		"Fixed the auth bug",
		s.Store.GetSession(a.UUID).Description,
	)
}

func TestEditSessionBoth(t *testing.T) {
	s := base.New(t)
	defer s.Close()
	a := s.NewSession(t)
	defer a.Close()
	a.Announce(a.Name(), "working")
	a.MustCallTool(
		constant.EditSession,
		map[string]any{
			constant.Alias:       "my-project",
			constant.Description: "Refactored the CLI",
		},
	)
	e := s.Store.GetSession(a.UUID)
	assert.String(t, "my-project", e.AliasValue())
	assert.String(t, "Refactored the CLI", e.Description)
}
