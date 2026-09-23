package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestGoAboveFloat(t *testing.T) {
	math_tester.AssertGoAboveFloat(t, false, 52, 51, 50)
	math_tester.AssertGoAboveFloat(t, false, 51, 50, 50)
	math_tester.AssertGoAboveFloat(t, false, 51, 49, 50)
	math_tester.AssertGoAboveFloat(t, true, 49, 51, 50)
}
