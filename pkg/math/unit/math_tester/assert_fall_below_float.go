package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/fall_below"
)

func (o *Tester) AssertFallBelowFloat(
	expected bool,
	past float64,
	now float64,
	threshold float64,
) {
	o.t.Helper()
	assert.Boolean(o.t, expected, fall_below.Float(past, now, threshold))
}
