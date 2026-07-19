package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestLinesAfter(t *testing.T) {
	assert.Strings(
		t,
		[]string{"c"},
		strings.LinesAfter([]string{"a", "b", "c"}, "b"),
	)
}
