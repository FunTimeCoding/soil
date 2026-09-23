package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/scale"
	"testing"
)

func AssertScaleFloat(
	t *testing.T,
	expected float64,
	from float64,
	to float64,
	factor float64,
) {
	t.Helper()
	assert.Round(t, expected, scale.Float(from, to, factor), 1)
}
