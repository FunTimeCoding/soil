package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/above_below"
	"testing"
)

func AssertAboveBelowFloat(
	t *testing.T,
	expected AboveBelow,
	f float64,
	magnitude float64,
) {
	t.Helper()
	var actual AboveBelow
	above_below.Float(
		f,
		magnitude,
		func() {
			actual.Above = true
		},
		func() {
			actual.Below = true
		},
	)
	assert.Any(t, expected, actual)
}
