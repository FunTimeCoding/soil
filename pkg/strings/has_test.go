package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestHas(t *testing.T) {
	assert.True(t, Has("abc", "a", "c"))
	assert.False(t, Has("abc", "a", "d"))
	assert.False(t, Has("abc", "x", "c"))
	assert.False(t, Has("abc", "x", "d"))
}
