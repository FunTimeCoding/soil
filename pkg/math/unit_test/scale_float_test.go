package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/scale"
	"testing"
)

func TestScaleFloat(t *testing.T) {
	scaleFloatAssertFloat(t, 0, 100, 0, 0)
	scaleFloatAssertFloat(t, 0, 100, 0.5, 50)
	scaleFloatAssertFloat(t, 0, 100, 1, 100)
	scaleFloatAssertFloat(t, 40, 20, 0.5, 30)
	scaleFloatAssertFloat(t, 1, 1, 0.5, 1)
	scaleFloatAssertFloat(t, -1, 1, 0.5, 0)
	scaleFloatAssertFloat(t, 2.0/3.0, 1, 0, 0.7)
	scaleFloatAssertFloat(t, 2.0/3.0, 1, 1, 1)
}

func scaleFloatAssertFloat(
	t *testing.T,
	from float64,
	to float64,
	factor float64,
	expect float64,
) {
	t.Helper()
	assert.Round(t, expect, scale.Float(from, to, factor), 1)
}
