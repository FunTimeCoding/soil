package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/scale"
)

func (o *Tester) AssertScaleInteger(
	expected int,
	from int,
	to int,
	factor float64,
) {
	o.t.Helper()
	assert.Integer(o.t, expected, scale.Value(from, to, factor))
}
