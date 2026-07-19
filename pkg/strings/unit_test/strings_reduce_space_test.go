package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestReduceSpace(t *testing.T) {
	assert.String(t, " ", strings.ReduceSpace("  "))
	assert.String(t, "a b", strings.ReduceSpace("a  b"))
	assert.String(t, "a b c", strings.ReduceSpace("a  b   c"))
}
