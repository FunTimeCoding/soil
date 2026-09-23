package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestGoAboveInteger(t *testing.T) {
	math_tester.AssertGoAboveInteger(t, false, 52, 51, 50)
	math_tester.AssertGoAboveInteger(t, false, 51, 50, 50)
	math_tester.AssertGoAboveInteger(t, false, 51, 49, 50)
	math_tester.AssertGoAboveInteger(t, true, 49, 51, 50)
}
