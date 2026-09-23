package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestNormalizeChangeFloat(t *testing.T) {
	math_tester.AssertNormalizeChangeFloat(t, 1, 0, 1, 0, 100)
	math_tester.AssertNormalizeChangeFloat(t, -1, 1, -2, 0, 100)
	math_tester.AssertNormalizeChangeFloat(t, -1, 1, -3, 0, 100)
	math_tester.AssertNormalizeChangeFloat(t, -100, 100, -100, 0, 100)
	math_tester.AssertNormalizeChangeFloat(t, -100, 100, -150, 0, 100)
	math_tester.AssertNormalizeChangeFloat(t, 5, 95, 20, 0, 100)
}

func TestFloatNoMaximum(t *testing.T) {
	math_tester.AssertNormalizeChangeFloat(t, 1, 0, 1, 0, 0)
	math_tester.AssertNormalizeChangeFloat(t, -1, 1, -2, 0, 0)
	math_tester.AssertNormalizeChangeFloat(t, -1, 1, -3, 0, 0)
	math_tester.AssertNormalizeChangeFloat(t, -100, 100, -100, 0, 0)
	math_tester.AssertNormalizeChangeFloat(t, -100, 100, -150, 0, 0)
}
