package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestReachIncreaseIntegerShortByOne(t *testing.T) {
	math_tester.AssertReachIncreaseInteger(t, false, 48, 49, 50)
}

func TestReachIncreaseIntegerReachedExactly(t *testing.T) {
	math_tester.AssertReachIncreaseInteger(t, true, 49, 50, 50)
}

func TestReachIncreaseIntegerExceedByOne(t *testing.T) {
	math_tester.AssertReachIncreaseInteger(t, true, 49, 51, 50)
}

func TestReachIncreaseIntegerDipBelow(t *testing.T) {
	math_tester.AssertReachIncreaseInteger(t, false, 51, 49, 50)
}
