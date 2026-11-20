package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/go_above"
)

func (o *Tester) AssertGoAboveInteger(
	expected bool,
	past int,
	now int,
	threshold int,
) {
	o.t.Helper()
	assert.Boolean(o.t, expected, go_above.Integer(past, now, threshold))
}
