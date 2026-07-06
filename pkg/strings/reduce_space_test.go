package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestReduceSpace(t *testing.T) {
	assert.String(t, " ", ReduceSpace("  "))
	assert.String(t, "a b", ReduceSpace("a  b"))
	assert.String(t, "a b c", ReduceSpace("a  b   c"))
}
