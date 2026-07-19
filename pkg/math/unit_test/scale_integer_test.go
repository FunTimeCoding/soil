package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/scale"
	"testing"
)

func TestScaleInteger(t *testing.T) {
	scaleIntegerAssertInteger(t, 0, 100, 0, 0)
	scaleIntegerAssertInteger(t, 0, 100, 0.5, 50)
	scaleIntegerAssertInteger(t, 0, 100, 1, 100)
	scaleIntegerAssertInteger(t, 40, 20, 0.5, 30)
	scaleIntegerAssertInteger(t, -10, -20, 0.5, -15)
	scaleIntegerAssertInteger(t, -30, -10, 0.5, -20)
}

func scaleIntegerAssertInteger(
	t *testing.T,
	from int,
	to int,
	factor float64,
	expect int,
) {
	t.Helper()
	assert.Integer(t, expect, scale.Integer(from, to, factor))
}
