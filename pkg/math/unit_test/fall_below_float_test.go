package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/fall_below"
	"testing"
)

func TestFallBelowFloat(t *testing.T) {
	// Short by 1
	fallBelowFloatAssertFloat(t, 51, 50, 50, false)
	// Reached exactly
	fallBelowFloatAssertFloat(t, 50, 49, 50, true)
	// Exceed by 1
	fallBelowFloatAssertFloat(t, 51, 49, 50, true)
	// Go above
	fallBelowFloatAssertFloat(t, 49, 51, 50, false)
}

func fallBelowFloatAssertFloat(
	t *testing.T,
	past float64,
	now float64,
	threshold float64,
	expect bool,
) {
	t.Helper()
	assert.Boolean(t, expect, fall_below.Float(past, now, threshold))
}
