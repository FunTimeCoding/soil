package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestReachReduceIntegerShortByOne(t *testing.T) {
	math_tester.AssertReachReduceInteger(t, false, 52, 51, 50)
}

func TestReachReduceIntegerReachedExactly(t *testing.T) {
	math_tester.AssertReachReduceInteger(t, true, 51, 50, 50)
}

func TestReachReduceIntegerExceedByOne(t *testing.T) {
	math_tester.AssertReachReduceInteger(t, true, 51, 49, 50)
}

func TestReachReduceIntegerRiseAbove(t *testing.T) {
	math_tester.AssertReachReduceInteger(t, false, 49, 51, 50)
}
