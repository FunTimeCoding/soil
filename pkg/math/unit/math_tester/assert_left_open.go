package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/in_range"
	"github.com/funtimecoding/soil/pkg/math/ranges"
	"testing"
)

func AssertLeftOpen(
	t *testing.T,
	expected bool,
	value float64,
	r ranges.Range,
) {
	t.Helper()
	assert.Boolean(t, expected, in_range.LeftOpen(value, r))
}
