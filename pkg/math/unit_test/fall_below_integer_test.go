package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/fall_below"
	"testing"
)

func TestFallBelowInteger(t *testing.T) {
	// Short by 1
	fallBelowIntegerAssertInteger(t, 51, 50, 50, false)
	// Reached exactly
	fallBelowIntegerAssertInteger(t, 50, 49, 50, true)
	// Exceed by 1
	fallBelowIntegerAssertInteger(t, 51, 49, 50, true)
	// Go above
	fallBelowIntegerAssertInteger(t, 49, 51, 50, false)
}

func fallBelowIntegerAssertInteger(
	t *testing.T,
	past int,
	now int,
	threshold int,
	expect bool,
) {
	t.Helper()
	assert.Boolean(t, expect, fall_below.Integer(past, now, threshold))
}
