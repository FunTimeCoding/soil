package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/normalize"
	"testing"
)

func TestNormalizeFloat(t *testing.T) {
	// Meet minimum
	normalizeFloatAssertFloat(t, 0, 0, 100, 0)
	// Below minimum
	normalizeFloatAssertFloat(t, -1, 0, 100, 0)
	// Meet maximum
	normalizeFloatAssertFloat(t, 100, 0, 100, 100)
	// Above maximum
	normalizeFloatAssertFloat(t, 101, 0, 100, 100)
	// No maximum
	normalizeFloatAssertFloat(t, 101, 0, 0, 101)
}

func normalizeFloatAssertFloat(
	t *testing.T,
	f float64,
	minimum float64,
	maximum float64,
	expect float64,
) {
	t.Helper()
	normalize.Float(&f, minimum, maximum)
	assert.Round(t, expect, f, 0)
}
