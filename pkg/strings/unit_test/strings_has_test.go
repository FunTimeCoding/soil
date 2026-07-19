package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestHas(t *testing.T) {
	assert.True(t, strings.Has("abc", "a", "c"))
	assert.False(t, strings.Has("abc", "a", "d"))
	assert.False(t, strings.Has("abc", "x", "c"))
	assert.False(t, strings.Has("abc", "x", "d"))
}
