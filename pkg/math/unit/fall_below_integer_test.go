package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestFallBelowIntegerShortByOne(t *testing.T) {
	math_tester.New(t).AssertFallBelowInteger(false, 51, 50, 50)
}

func TestFallBelowIntegerReachedExactly(t *testing.T) {
	math_tester.New(t).AssertFallBelowInteger(true, 50, 49, 50)
}

func TestFallBelowIntegerExceedByOne(t *testing.T) {
	math_tester.New(t).AssertFallBelowInteger(true, 51, 49, 50)
}

func TestFallBelowIntegerGoAbove(t *testing.T) {
	math_tester.New(t).AssertFallBelowInteger(false, 49, 51, 50)
}
