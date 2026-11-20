package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math"
)

func (o *Tester) AssertExceededByPercent(
	expected bool,
	past float64,
	present float64,
	hundredPercent float64,
	percent float64,
) {
	o.t.Helper()
	assert.Boolean(
		o.t,
		expected,
		math.ExceededByPercent(past, present, hundredPercent, percent),
	)
}
