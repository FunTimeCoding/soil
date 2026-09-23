package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestWeight(t *testing.T) {
	math_tester.AssertWeight(t, 1, 1, 1, 1, 1)
	math_tester.AssertWeight(t, 0.75, 1, 0.5, 1, 1)
	math_tester.AssertWeight(t, 0.5, 1, 0, 1, 1)
	math_tester.AssertWeight(t, 127.5, 255, 0, 1, 1)
	math_tester.AssertWeight(t, 0.67, 1, 0, 2, 1)
	math_tester.AssertWeight(t, 0.33, 1, 0, 1, 2)
}
