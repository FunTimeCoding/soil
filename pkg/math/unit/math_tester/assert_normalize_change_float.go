package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/normalize_change"
	"testing"
)

func AssertNormalizeChangeFloat(
	t *testing.T,
	expected float64,
	now float64,
	change float64,
	minimum float64,
	maximum float64,
) {
	t.Helper()
	assert.Float(
		t,
		expected,
		normalize_change.Float(now, change, minimum, maximum),
	)
}
