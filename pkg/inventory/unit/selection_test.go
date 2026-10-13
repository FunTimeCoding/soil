package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/inventory/selection"
	"testing"
)

func TestSelectionRemembersPerSession(t *testing.T) {
	s := selection.New()
	s.Set("one", "alfa")
	s.Set("two", "bravo")
	active, okay := s.Active("one")
	assert.True(t, okay)
	assert.String(t, "alfa", active)
	other, _ := s.Active("two")
	assert.String(t, "bravo", other)
}

func TestSelectionClearForgetsTheSession(t *testing.T) {
	s := selection.New()
	s.Set("one", "alfa")
	s.Clear("one")
	_, okay := s.Active("one")
	assert.False(t, okay)
}
