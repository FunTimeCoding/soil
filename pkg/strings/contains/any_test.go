package contains

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestAny(t *testing.T) {
	assert.True(t, Any([]string{upper.Alfa}, []string{upper.Alfa}))
	assert.False(t, Any([]string{upper.Alfa}, []string{upper.Bravo}))
	assert.True(
		t,
		Any(
			[]string{upper.Alfa, upper.Bravo, upper.Charlie},
			[]string{upper.Alfa},
		),
	)
	assert.False(
		t,
		Any(
			[]string{upper.Alfa, upper.Bravo, upper.Charlie},
			[]string{upper.Delta},
		),
	)
}
