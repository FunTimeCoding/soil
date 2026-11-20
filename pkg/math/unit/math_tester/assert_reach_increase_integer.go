package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/reach_increase"
)

func (o *Tester) AssertReachIncreaseInteger(
	expected bool,
	past int,
	now int,
	threshold int,
) {
	o.t.Helper()
	assert.Boolean(o.t, expected, reach_increase.Integer(past, now, threshold))
}
