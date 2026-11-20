package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/in_range"
	"github.com/funtimecoding/soil/pkg/math/ranges"
)

func (o *Tester) AssertRightOpen(
	expected bool,
	value float64,
	r ranges.Range,
) {
	o.t.Helper()
	assert.Boolean(o.t, expected, in_range.RightOpen(value, r))
}
