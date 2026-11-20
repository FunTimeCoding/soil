package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/normalize"
)

func (o *Tester) AssertNormalizeFloat(
	expected float64,
	f float64,
	minimum float64,
	maximum float64,
) {
	o.t.Helper()
	normalize.Float(&f, minimum, maximum)
	assert.Round(o.t, expected, f, 0)
}
