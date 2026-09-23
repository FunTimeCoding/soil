package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math"
	"testing"
)

func AssertSeriesSum(
	t *testing.T,
	expected int,
	n int,
) {
	t.Helper()
	assert.Integer(t, expected, math.SeriesSum(n))
}
