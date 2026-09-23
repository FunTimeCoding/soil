package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestReachReduceFloatShortByOne(t *testing.T) {
	math_tester.AssertReachReduceFloat(t, false, 52, 51, 50)
}

func TestReachReduceFloatReachedExactly(t *testing.T) {
	math_tester.AssertReachReduceFloat(t, true, 51, 50, 50)
}

func TestReachReduceFloatExceedByOne(t *testing.T) {
	math_tester.AssertReachReduceFloat(t, true, 51, 49, 50)
}

func TestReachReduceFloatRiseAbove(t *testing.T) {
	math_tester.AssertReachReduceFloat(t, false, 49, 51, 50)
}
