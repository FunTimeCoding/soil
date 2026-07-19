package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestIndent(t *testing.T) {
	assert.String(
		t,
		"    hello",
		strings.Indent("hello", 4, false),
	)
	assert.String(
		t,
		"    A\n    \n    B",
		strings.Indent("A\n\nB", 4, false),
	)
	assert.String(
		t,
		"    A\n    B",
		strings.Indent("A\n\nB", 4, true),
	)
}
