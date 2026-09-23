package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/reach_reduce"
	"testing"
)

func AssertReachReduceInteger(
	t *testing.T,
	expected bool,
	past int,
	now int,
	threshold int,
) {
	t.Helper()
	assert.Boolean(t, expected, reach_reduce.Integer(past, now, threshold))
}
