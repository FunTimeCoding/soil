package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestElementAfter(t *testing.T) {
	assert.String(
		t,
		"b",
		strings.ElementAfter([]string{"a", "b"}, "a"),
	)
}
