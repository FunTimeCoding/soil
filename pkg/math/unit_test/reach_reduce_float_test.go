package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/reach_reduce"
	"testing"
)

func TestReachReduceFloat(t *testing.T) {
	// Short by 1
	reachReduceFloatAssertFloat(t, 52, 51, 50, false)
	// Reached exactly
	reachReduceFloatAssertFloat(t, 51, 50, 50, true)
	// Exceed by 1
	reachReduceFloatAssertFloat(t, 51, 49, 50, true)
	// Reach increase
	reachReduceFloatAssertFloat(t, 49, 51, 50, false)
}

func reachReduceFloatAssertFloat(
	t *testing.T,
	past float64,
	now float64,
	threshold float64,
	expect bool,
) {
	t.Helper()
	assert.Boolean(t, expect, reach_reduce.Float(past, now, threshold))
}
