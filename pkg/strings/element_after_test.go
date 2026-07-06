package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestElementAfter(t *testing.T) {
	assert.String(
		t,
		"b",
		ElementAfter([]string{"a", "b"}, "a"),
	)
}
