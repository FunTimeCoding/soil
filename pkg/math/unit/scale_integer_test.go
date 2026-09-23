package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestScaleInteger(t *testing.T) {
	math_tester.AssertScaleInteger(t, 0, 0, 100, 0)
	math_tester.AssertScaleInteger(t, 50, 0, 100, 0.5)
	math_tester.AssertScaleInteger(t, 100, 0, 100, 1)
	math_tester.AssertScaleInteger(t, 30, 40, 20, 0.5)
	math_tester.AssertScaleInteger(t, -15, -10, -20, 0.5)
	math_tester.AssertScaleInteger(t, -20, -30, -10, 0.5)
}
