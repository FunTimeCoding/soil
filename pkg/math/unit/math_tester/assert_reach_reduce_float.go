package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/reach_reduce"
)

func (o *Tester) AssertReachReduceFloat(
	expected bool,
	past float64,
	now float64,
	threshold float64,
) {
	o.t.Helper()
	assert.Boolean(o.t, expected, reach_reduce.Float(past, now, threshold))
}
