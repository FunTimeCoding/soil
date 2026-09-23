package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/reach_reduce"
	"testing"
)

func AssertReachReduceFloat(
	t *testing.T,
	expected bool,
	past float64,
	now float64,
	threshold float64,
) {
	t.Helper()
	assert.Boolean(t, expected, reach_reduce.Float(past, now, threshold))
}
