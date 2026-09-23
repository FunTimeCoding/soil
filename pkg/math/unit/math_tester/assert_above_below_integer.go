package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/above_below"
	"testing"
)

func AssertAboveBelowInteger(
	t *testing.T,
	expected AboveBelow,
	f int,
	magnitude int,
) {
	t.Helper()
	var actual AboveBelow
	above_below.Integer(
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
