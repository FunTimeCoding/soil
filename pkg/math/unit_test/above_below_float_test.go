package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/above_below"
	"testing"
)

func TestAboveBelowFloat(t *testing.T) {
	aboveBelowFloatAssertFloat(t, 1, 0, true, false)
	aboveBelowFloatAssertFloat(t, -1, 0, false, true)
	aboveBelowFloatAssertFloat(t, 0, 0, false, false)
	aboveBelowFloatAssertFloat(t, 1, 1, false, false)
	aboveBelowFloatAssertFloat(t, -1, 1, false, false)
	aboveBelowFloatAssertFloat(t, 0, 1, false, false)
	aboveBelowFloatAssertFloat(t, 2, 1, true, false)
	aboveBelowFloatAssertFloat(t, -2, 1, false, true)
}

func aboveBelowFloatAssertFloat(
	t *testing.T,
	f float64,
	magnitude float64,
	expectAbove bool,
	expectBelow bool,
) {
	t.Helper()
	var above bool
	var below bool
	above_below.Float(
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
