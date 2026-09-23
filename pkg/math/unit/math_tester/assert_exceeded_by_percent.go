package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math"
	"testing"
)

func AssertExceededByPercent(
	t *testing.T,
	expected bool,
	past float64,
	present float64,
	hundredPercent float64,
	percent float64,
) {
	t.Helper()
	assert.Boolean(
		t,
		expected,
		math.ExceededByPercent(past, present, hundredPercent, percent),
	)
}
