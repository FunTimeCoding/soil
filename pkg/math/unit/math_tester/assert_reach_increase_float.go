package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/reach_increase"
)

func (o *Tester) AssertReachIncreaseFloat(
	expected bool,
	past float64,
	now float64,
	threshold float64,
) {
	o.t.Helper()
	assert.Boolean(o.t, expected, reach_increase.Float(past, now, threshold))
}
