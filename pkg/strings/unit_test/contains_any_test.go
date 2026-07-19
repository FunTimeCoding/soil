package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/contains"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestAny(t *testing.T) {
	assert.True(t, contains.Any([]string{upper.Alfa}, []string{upper.Alfa}))
	assert.False(t, contains.Any([]string{upper.Alfa}, []string{upper.Bravo}))
	assert.True(
		t,
		contains.Any(
			[]string{upper.Alfa, upper.Bravo, upper.Charlie},
			[]string{upper.Alfa},
		),
	)
	assert.False(
		t,
		contains.Any(
			[]string{upper.Alfa, upper.Bravo, upper.Charlie},
			[]string{upper.Delta},
		),
	)
}
