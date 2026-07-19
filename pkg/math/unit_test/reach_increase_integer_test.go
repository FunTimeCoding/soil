package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/reach_increase"
	"testing"
)

func TestReachIncreaseInteger(t *testing.T) {
	// Short by 1
	reachIncreaseIntegerAssertInteger(t, 48, 49, 50, false)
	// Reached exactly
	reachIncreaseIntegerAssertInteger(t, 49, 50, 50, true)
	// Exceed by 1
	reachIncreaseIntegerAssertInteger(t, 49, 51, 50, true)
	// Dip below
	reachIncreaseIntegerAssertInteger(t, 51, 49, 50, false)
}

func reachIncreaseIntegerAssertInteger(
	t *testing.T,
	past int,
	now int,
	threshold int,
	expect bool,
) {
	t.Helper()
	assert.Boolean(t, expect, reach_increase.Integer(past, now, threshold))
}
