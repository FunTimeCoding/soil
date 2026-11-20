package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/normalize_change"
)

func (o *Tester) AssertNormalizeChangeFloat(
	expected float64,
	now float64,
	change float64,
	minimum float64,
	maximum float64,
) {
	o.t.Helper()
	assert.Float(
		o.t,
		expected,
		normalize_change.Value(now, change, minimum, maximum),
	)
}
