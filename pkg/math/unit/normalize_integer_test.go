package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestNormalizeIntegerMeetMinimum(t *testing.T) {
	math_tester.AssertNormalizeInteger(t, 0, 0, 0, 100)
}

func TestNormalizeIntegerBelowMinimum(t *testing.T) {
	math_tester.AssertNormalizeInteger(t, 0, -1, 0, 100)
}

func TestNormalizeIntegerMeetMaximum(t *testing.T) {
	math_tester.AssertNormalizeInteger(t, 100, 100, 0, 100)
}

func TestNormalizeIntegerAboveMaximum(t *testing.T) {
	math_tester.AssertNormalizeInteger(t, 100, 101, 0, 100)
}

func TestNormalizeIntegerNoMaximum(t *testing.T) {
	math_tester.AssertNormalizeInteger(t, 101, 101, 0, 0)
}
