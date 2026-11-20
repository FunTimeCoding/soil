package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/above_below"
)

func (o *Tester) AssertAboveBelowInteger(
	expected AboveBelow,
	f int,
	magnitude int,
) {
	o.t.Helper()
	var actual AboveBelow
	above_below.Call(
		f,
		magnitude,
		func() {
			actual.Above = true
		},
		func() {
			actual.Below = true
		},
	)
	assert.Any(o.t, expected, actual)
}
