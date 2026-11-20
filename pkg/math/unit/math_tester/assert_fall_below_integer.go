package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/fall_below"
)

func (o *Tester) AssertFallBelowInteger(
	expected bool,
	past int,
	now int,
	threshold int,
) {
	o.t.Helper()
	assert.Boolean(o.t, expected, fall_below.Integer(past, now, threshold))
}
