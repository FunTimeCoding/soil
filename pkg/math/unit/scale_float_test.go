package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestScaleFloat(t *testing.T) {
	math_tester.AssertScaleFloat(t, 0, 0, 100, 0)
	math_tester.AssertScaleFloat(t, 50, 0, 100, 0.5)
	math_tester.AssertScaleFloat(t, 100, 0, 100, 1)
	math_tester.AssertScaleFloat(t, 30, 40, 20, 0.5)
	math_tester.AssertScaleFloat(t, 1, 1, 1, 0.5)
	math_tester.AssertScaleFloat(t, 0, -1, 1, 0.5)
	math_tester.AssertScaleFloat(t, 0.7, 2.0/3.0, 1, 0)
	math_tester.AssertScaleFloat(t, 1, 2.0/3.0, 1, 1)
}
