package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestSecondHalf(t *testing.T) {
	assert.Any(
		t,
		[]string{"Charlie", "Delta"},
		SecondHalf(
			[]string{
				upper.Alfa,
				upper.Bravo,
				upper.Charlie,
				upper.Delta,
			},
		),
	)
}
