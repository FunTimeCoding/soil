package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/scale"
)

func (o *Tester) AssertScaleFloat(
	expected float64,
	from float64,
	to float64,
	factor float64,
) {
	o.t.Helper()
	assert.Round(o.t, expected, scale.Value(from, to, factor), 1)
}
