package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/normalize_change"
)

func (o *Tester) AssertNormalizeChangeInteger(
	expected int,
	now int,
	change int,
	minimum int,
	maximum int,
) {
	o.t.Helper()
	assert.Integer(
		o.t,
		expected,
		normalize_change.Value(now, change, minimum, maximum),
	)
}
