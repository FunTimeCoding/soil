package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/join"
	"testing"
)

func TestRelative(t *testing.T) {
	assert.String(t, "a/b/c", join.Relative("a", "b", "c"))
	assert.String(t, "a/b/c", join.Relative("/a", "b", "c"))
	assert.String(t, "a/b/c", join.Relative("a", "/b", "c"))
	assert.String(t, "a/b/c", join.Relative("a", "b", "/c"))
	assert.String(t, "a/b/c", join.Relative("/a", "/b", "c"))
	assert.String(t, "a/b/c", join.Relative("a", "/b", "/c"))
	assert.String(t, "a/b/c", join.Relative("/a", "b", "/c"))
	assert.String(t, "a/b/c", join.Relative("/a/", "/b/", "/c/"))
}
