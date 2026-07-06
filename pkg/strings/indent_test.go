package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestIndent(t *testing.T) {
	assert.String(
		t,
		"    hello",
		Indent("hello", 4, false),
	)
	assert.String(
		t,
		"    A\n    \n    B",
		Indent("A\n\nB", 4, false),
	)
	assert.String(
		t,
		"    A\n    B",
		Indent("A\n\nB", 4, true),
	)
}
