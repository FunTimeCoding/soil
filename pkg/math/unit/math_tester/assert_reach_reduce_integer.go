package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/reach_reduce"
)

func (o *Tester) AssertReachReduceInteger(
	expected bool,
	past int,
	now int,
	threshold int,
) {
	o.t.Helper()
	assert.Boolean(o.t, expected, reach_reduce.Integer(past, now, threshold))
}
