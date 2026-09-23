package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math"
	"testing"
)

func AssertWeight(
	t *testing.T,
	expected float64,
	a float64,
	b float64,
	aWeight float64,
	bWeight float64,
) {
	t.Helper()
	assert.Round(t, expected, math.Weight(a, b, aWeight, bWeight), 2)
}
