package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestSeriesSum(t *testing.T) {
	math_tester.AssertSeriesSum(t, 1, 1)
	math_tester.AssertSeriesSum(t, 45, 9)
}
