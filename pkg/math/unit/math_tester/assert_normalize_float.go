package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/normalize"
	"testing"
)

func AssertNormalizeFloat(
	t *testing.T,
	expected float64,
	f float64,
	minimum float64,
	maximum float64,
) {
	t.Helper()
	normalize.Float(&f, minimum, maximum)
	assert.Round(t, expected, f, 0)
}
