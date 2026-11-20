package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestReachIncreaseFloatShortByOne(t *testing.T) {
	math_tester.New(t).AssertReachIncreaseFloat(false, 48, 49, 50)
}

func TestReachIncreaseFloatReachedExactly(t *testing.T) {
	math_tester.New(t).AssertReachIncreaseFloat(true, 49, 50, 50)
}

func TestReachIncreaseFloatExceedByOne(t *testing.T) {
	math_tester.New(t).AssertReachIncreaseFloat(true, 49, 51, 50)
}

func TestReachIncreaseFloatDipBelow(t *testing.T) {
	math_tester.New(t).AssertReachIncreaseFloat(false, 51, 49, 50)
}
