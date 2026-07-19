package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/normalize_change"
	"testing"
)

func TestNormalizeChangeInteger(t *testing.T) {
	normalizeChangeIntegerAssertInteger(t, 0, 1, 0, 100, 1)
	normalizeChangeIntegerAssertInteger(t, 1, -2, 0, 100, -1)
	normalizeChangeIntegerAssertInteger(t, 1, -3, 0, 100, -1)
	normalizeChangeIntegerAssertInteger(t, 100, -100, 0, 100, -100)
	normalizeChangeIntegerAssertInteger(t, 100, -150, 0, 100, -100)
	normalizeChangeIntegerAssertInteger(t, 95, 20, 0, 100, 5)
}

func TestIntegerNoMaximum(t *testing.T) {
	normalizeChangeIntegerAssertInteger(t, 0, 1, 0, 0, 1)
	normalizeChangeIntegerAssertInteger(t, 1, -2, 0, 0, -1)
	normalizeChangeIntegerAssertInteger(t, 1, -3, 0, 0, -1)
	normalizeChangeIntegerAssertInteger(t, 100, -100, 0, 0, -100)
	normalizeChangeIntegerAssertInteger(t, 100, -150, 0, 0, -100)
}

func normalizeChangeIntegerAssertInteger(
	t *testing.T,
	now int,
	change int,
	minimum int,
	maximum int,
	expect int,
) {
	t.Helper()
	assert.Integer(
		t,
		expect,
		normalize_change.Integer(
			now,
			change,
			minimum,
			maximum,
		),
	)
}
