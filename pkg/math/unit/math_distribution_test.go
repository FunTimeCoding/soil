package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestDistribution(t *testing.T) {
	math_tester.AssertDistribution(t, 100, 100, 10)
	math_tester.AssertDistribution(t, 1000000, 1000000, 10)
}
