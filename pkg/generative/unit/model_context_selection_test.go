package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/model_context/selection"
	"github.com/funtimecoding/soil/pkg/generative/unit/session_tester"
	"github.com/mark3labs/mcp-go/server"
	"testing"
)

func TestSelectionRemembersPerSession(t *testing.T) {
	m := server.NewMCPServer("test", "0")
	one := m.WithContext(t.Context(), session_tester.New("one"))
	two := m.WithContext(t.Context(), session_tester.New("two"))
	s := selection.New()
	assert.FatalOnError(t, s.Select(one, "delta"))
	assert.FatalOnError(t, s.Select(two, "bravo"))
	active, e := s.Active(one)
	assert.FatalOnError(t, e)
	assert.String(t, "delta", active)
	other, f := s.Active(two)
	assert.FatalOnError(t, f)
	assert.String(t, "bravo", other)
}

func TestSelectionIsEmptyUntilChosen(t *testing.T) {
	m := server.NewMCPServer("test", "0")
	x := m.WithContext(t.Context(), session_tester.New("one"))
	active, e := selection.New().Active(x)
	assert.FatalOnError(t, e)
	assert.String(t, "", active)
}

func TestSelectionClearForgetsTheSession(t *testing.T) {
	m := server.NewMCPServer("test", "0")
	x := m.WithContext(t.Context(), session_tester.New("one"))
	s := selection.New()
	assert.FatalOnError(t, s.Select(x, "delta"))
	s.Clear("one")
	active, e := s.Active(x)
	assert.FatalOnError(t, e)
	assert.String(t, "", active)
}

func TestSelectionRefusesAContextWithoutASession(t *testing.T) {
	s := selection.New()
	assert.NotNil(t, s.Select(t.Context(), "delta"))
	_, e := s.Active(t.Context())
	assert.StringContains(t, "no model context session", e.Error())
}
