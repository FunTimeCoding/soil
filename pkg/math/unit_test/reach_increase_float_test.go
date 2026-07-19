package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/reach_increase"
	"testing"
)

func TestReachIncreaseFloat(t *testing.T) {
	// Short by 1
	reachIncreaseFloatAssertFloat(t, 48, 49, 50, false)
	// Reached exactly
	reachIncreaseFloatAssertFloat(t, 49, 50, 50, true)
	// Exceed by 1
	reachIncreaseFloatAssertFloat(t, 49, 51, 50, true)
	// Reach decrease
	reachIncreaseFloatAssertFloat(t, 51, 49, 50, false)
}

func reachIncreaseFloatAssertFloat(
	t *testing.T,
	past float64,
	now float64,
	threshold float64,
	expect bool,
) {
	t.Helper()
	assert.Boolean(t, expect, reach_increase.Float(past, now, threshold))
}
