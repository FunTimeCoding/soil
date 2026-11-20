package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math"
)

func (o *Tester) AssertSeriesSum(
	expected int,
	n int,
) {
	o.t.Helper()
	assert.Integer(o.t, expected, math.SeriesSum(n))
}
