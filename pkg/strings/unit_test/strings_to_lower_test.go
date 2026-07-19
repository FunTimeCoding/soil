package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestToLower(t *testing.T) {
	assert.Strings(
		t,
		[]string{"alfa", "bravo"},
		strings.ToLower([]string{upper.Alfa, upper.Bravo}),
	)
}
