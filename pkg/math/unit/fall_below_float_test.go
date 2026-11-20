package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestFallBelowFloatShortByOne(t *testing.T) {
	math_tester.New(t).AssertFallBelowFloat(false, 51, 50, 50)
}

func TestFallBelowFloatReachedExactly(t *testing.T) {
	math_tester.New(t).AssertFallBelowFloat(true, 50, 49, 50)
}

func TestFallBelowFloatExceedByOne(t *testing.T) {
	math_tester.New(t).AssertFallBelowFloat(true, 51, 49, 50)
}

func TestFallBelowFloatGoAbove(t *testing.T) {
	math_tester.New(t).AssertFallBelowFloat(false, 49, 51, 50)
}
