package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/scale"
	"testing"
)

func AssertScaleInteger(
	t *testing.T,
	expected int,
	from int,
	to int,
	factor float64,
) {
	t.Helper()
	assert.Integer(t, expected, scale.Integer(from, to, factor))
}
