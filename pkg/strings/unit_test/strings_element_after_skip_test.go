package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestElementAfterSkip(t *testing.T) {
	assert.String(
		t,
		"b",
		strings.ElementAfterSkip([]string{"a", "b", "a", "c"}, "a", 0),
	)
	assert.String(
		t,
		"c",
		strings.ElementAfterSkip([]string{"a", "b", "a", "c"}, "a", 1),
	)
}
