package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/go_above"
)

func (o *Tester) AssertGoAboveFloat(
	expected bool,
	past float64,
	now float64,
	threshold float64,
) {
	o.t.Helper()
	assert.Boolean(o.t, expected, go_above.Float(past, now, threshold))
}
