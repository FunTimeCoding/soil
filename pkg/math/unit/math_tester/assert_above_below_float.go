package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/above_below"
)

func (o *Tester) AssertAboveBelowFloat(
	expected AboveBelow,
	f float64,
	magnitude float64,
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
