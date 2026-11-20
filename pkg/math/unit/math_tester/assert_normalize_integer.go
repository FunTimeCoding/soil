package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/normalize"
)

func (o *Tester) AssertNormalizeInteger(
	expected int,
	i int,
	minimum int,
	maximum int,
) {
	o.t.Helper()
	normalize.Integer(&i, minimum, maximum)
	assert.Integer(o.t, expected, i)
}
