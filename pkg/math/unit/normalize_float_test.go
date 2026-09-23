package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestNormalizeFloatMeetMinimum(t *testing.T) {
	math_tester.AssertNormalizeFloat(t, 0, 0, 0, 100)
}

func TestNormalizeFloatBelowMinimum(t *testing.T) {
	math_tester.AssertNormalizeFloat(t, 0, -1, 0, 100)
}

func TestNormalizeFloatMeetMaximum(t *testing.T) {
	math_tester.AssertNormalizeFloat(t, 100, 100, 0, 100)
}

func TestNormalizeFloatAboveMaximum(t *testing.T) {
	math_tester.AssertNormalizeFloat(t, 100, 101, 0, 100)
}

func TestNormalizeFloatNoMaximum(t *testing.T) {
	math_tester.AssertNormalizeFloat(t, 101, 101, 0, 0)
}
