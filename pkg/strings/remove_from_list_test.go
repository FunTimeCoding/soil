package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestRemoveFromList(t *testing.T) {
	assert.Any(
		t,
		[]string{upper.Bravo, upper.Charlie},
		RemoveFromList(
			[]string{upper.Alfa, upper.Alfa, upper.Bravo, upper.Charlie},
			[]string{upper.Alfa},
		),
	)
	assert.Any(
		t,
		[]string{upper.Charlie},
		RemoveFromList(
			[]string{upper.Alfa, upper.Alfa, upper.Bravo, upper.Charlie},
			[]string{upper.Alfa, upper.Bravo},
		),
	)
}
