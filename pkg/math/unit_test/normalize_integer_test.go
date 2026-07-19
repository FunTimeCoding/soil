package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/normalize"
	"testing"
)

func TestNormalizeInteger(t *testing.T) {
	// Meet minimum
	normalizeIntegerAssertInteger(t, 0, 0, 100, 0)
	// Below minimum
	normalizeIntegerAssertInteger(t, -1, 0, 100, 0)
	// Meet maximum
	normalizeIntegerAssertInteger(t, 100, 0, 100, 100)
	// Above maximum
	normalizeIntegerAssertInteger(t, 101, 0, 100, 100)
	// No maximum
	normalizeIntegerAssertInteger(t, 101, 0, 0, 101)
}

func normalizeIntegerAssertInteger(
	t *testing.T,
	i int,
	minimum int,
	maximum int,
	expect int,
) {
	t.Helper()
	normalize.Integer(&i, minimum, maximum)
	assert.Integer(t, expect, i)
}
