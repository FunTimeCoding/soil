package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/not_selected"
	"github.com/funtimecoding/soil/pkg/inventory"
	"testing"
)

func TestResolveExplicit(t *testing.T) {
	name, e := inventory.Resolve("second", []string{"first", "second"})
	assert.Nil(t, e)
	assert.String(t, "second", name)
}

func TestResolveUnknownIsNotFound(t *testing.T) {
	_, e := inventory.Resolve("third", []string{"first", "second"})
	assert.True(t, not_found.Is(e))
	assert.String(t, "instance not found: third", e.Error())
}

func TestResolveSingleAutoSelects(t *testing.T) {
	name, e := inventory.Resolve("", []string{"only"})
	assert.Nil(t, e)
	assert.String(t, "only", name)
}

func TestResolveManyIsNotSelected(t *testing.T) {
	_, e := inventory.Resolve("", []string{"first", "second"})
	assert.True(t, not_selected.Is(e))
	assert.String(
		t,
		"no instance selected - 2 instances configured, selection required",
		e.Error(),
	)
}
