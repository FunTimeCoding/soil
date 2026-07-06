package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestLinesAfter(t *testing.T) {
	assert.Strings(
		t,
		[]string{"c"},
		LinesAfter([]string{"a", "b", "c"}, "b"),
	)
}
