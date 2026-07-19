package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/above_below"
	"testing"
)

func TestAboveBelowInteger(t *testing.T) {
	aboveBelowIntegerAssertInteger(t, 1, 0, true, false)
	aboveBelowIntegerAssertInteger(t, -1, 0, false, true)
	aboveBelowIntegerAssertInteger(t, 0, 0, false, false)
}

func aboveBelowIntegerAssertInteger(
	t *testing.T,
	f int,
	magnitude int,
	expectAbove bool,
	expectBelow bool,
) {
	t.Helper()
	var above bool
	var below bool
	above_below.Integer(
		f,
		magnitude,
		func() {
			above = true
		},
		func() {
			below = true
		},
	)
	assert.Boolean(t, expectAbove, above)
	assert.Boolean(t, expectBelow, below)
}
