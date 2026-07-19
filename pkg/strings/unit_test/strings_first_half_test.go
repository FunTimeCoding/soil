package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestFirstHalf(t *testing.T) {
	assert.Any(
		t,
		[]string{"Alfa", "Bravo"},
		strings.FirstHalf(
			[]string{
				upper.Alfa,
				upper.Bravo,
				upper.Charlie,
				upper.Delta,
			},
		),
	)
}
