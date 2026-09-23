package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestFallBelowIntegerShortByOne(t *testing.T) {
	math_tester.AssertFallBelowInteger(t, false, 51, 50, 50)
}

func TestFallBelowIntegerReachedExactly(t *testing.T) {
	math_tester.AssertFallBelowInteger(t, true, 50, 49, 50)
}

func TestFallBelowIntegerExceedByOne(t *testing.T) {
	math_tester.AssertFallBelowInteger(t, true, 51, 49, 50)
}

func TestFallBelowIntegerGoAbove(t *testing.T) {
	math_tester.AssertFallBelowInteger(t, false, 49, 51, 50)
}
