package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestElementAfterSkip(t *testing.T) {
	assert.String(
		t,
		"b",
		ElementAfterSkip([]string{"a", "b", "a", "c"}, "a", 0),
	)
	assert.String(
		t,
		"c",
		ElementAfterSkip([]string{"a", "b", "a", "c"}, "a", 1),
	)
}
