package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/fall_below"
	"testing"
)

func AssertFallBelowFloat(
	t *testing.T,
	expected bool,
	past float64,
	now float64,
	threshold float64,
) {
	t.Helper()
	assert.Boolean(t, expected, fall_below.Float(past, now, threshold))
}
