package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/normalize_change"
	"testing"
)

func TestNormalizeChangeFloat(t *testing.T) {
	normalizeChangeFloatAssertFloat(t, 0, 1, 0, 100, 1)
	normalizeChangeFloatAssertFloat(t, 1, -2, 0, 100, -1)
	normalizeChangeFloatAssertFloat(t, 1, -3, 0, 100, -1)
	normalizeChangeFloatAssertFloat(t, 100, -100, 0, 100, -100)
	normalizeChangeFloatAssertFloat(t, 100, -150, 0, 100, -100)
	normalizeChangeFloatAssertFloat(t, 95, 20, 0, 100, 5)
}

func TestFloatNoMaximum(t *testing.T) {
	normalizeChangeFloatAssertFloat(t, 0, 1, 0, 0, 1)
	normalizeChangeFloatAssertFloat(t, 1, -2, 0, 0, -1)
	normalizeChangeFloatAssertFloat(t, 1, -3, 0, 0, -1)
	normalizeChangeFloatAssertFloat(t, 100, -100, 0, 0, -100)
	normalizeChangeFloatAssertFloat(t, 100, -150, 0, 0, -100)
}

func normalizeChangeFloatAssertFloat(
	t *testing.T,
	now float64,
	change float64,
	minimum float64,
	maximum float64,
	expect float64,
) {
	t.Helper()
	assert.Float(t, expect, normalize_change.Float(now, change, minimum, maximum))
}
