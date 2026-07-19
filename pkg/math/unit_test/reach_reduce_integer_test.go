package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/reach_reduce"
	"testing"
)

func TestReachReduceInteger(t *testing.T) {
	// Short by 1
	reachReduceIntegerAssertInteger(t, 52, 51, 50, false)
	// Reached exactly
	reachReduceIntegerAssertInteger(t, 51, 50, 50, true)
	// Exceed by 1
	reachReduceIntegerAssertInteger(t, 51, 49, 50, true)
	// Reach increase
	reachReduceIntegerAssertInteger(t, 49, 51, 50, false)
}

func reachReduceIntegerAssertInteger(
	t *testing.T,
	past int,
	now int,
	threshold int,
	expect bool,
) {
	t.Helper()
	assert.Boolean(t, expect, reach_reduce.Integer(past, now, threshold))
}
